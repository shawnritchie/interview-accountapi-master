package form3

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	CREATE 				= "CREATE"
	FETCH				= "FETCH"
	LIST 				= "LIST"
	DELETE 				= "DELETE"

	F3BaseURL           = "F3BaseURL"
	F3Timeout           = "F3Timeout"
	F3MaxRetries        = "F3MaxRetries"
)

type F3Env struct {
	F3BaseURL    string
	F3Timeout    time.Duration
	F3MaxRetries int
}

var clientEnv F3Env
var Logger *log.Logger

func init() {
	Logger = log.New(os.Stderr, "F3CLIENT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type F3Client struct {
	Env        F3Env
	HTTPClient *http.Client
}

func SetupF3Client(env F3Env) *F3Client {
	return &F3Client{
		Env: env,
		HTTPClient: &http.Client{
			Timeout: clientEnv.F3Timeout,
		},
	}
}

func NewF3Client() (*F3Client, error) {
	clientEnv.F3BaseURL = os.Getenv(F3BaseURL)
	if clientEnv.F3BaseURL == "" {
		return nil, fmt.Errorf("'F3BaseURL' environmental variable not set")
	}

	if duration, err := time.ParseDuration(os.Getenv(F3Timeout)); err == nil {
		clientEnv.F3Timeout = duration
	} else {
		Logger.Printf("F3Timeout environment variable not set. defaulting to 60 seconds")
		clientEnv.F3Timeout = 60 * time.Second
	}

	if retries, err := strconv.Atoi(os.Getenv(F3MaxRetries)); err == nil {
		clientEnv.F3MaxRetries = retries
	} else {
		Logger.Printf("F3MaxRetries environment variable not set. defaulting to 3 retries")
		clientEnv.F3MaxRetries = 3
	}

	return SetupF3Client(clientEnv), nil
}

func (c *F3Client) Create() CreateBuilder {
	return newAccountBuilder(c)
}

func (c *F3Client) Fetch() FetchBuilder {
	return newFetchBuilder(c)
}

func (c *F3Client) List() ListBuilder {
	return newListBuilder(c)
}

func (c *F3Client) Delete() DeleteBuilder {
	return newDeleteBuilder(c)
}

func (c *F3Client) request(req *http.Request, body interface{}) error {
	req.Header.Set("Host", c.Env.F3BaseURL)
	req.Header.Set("Date", time.Now().Format(time.RFC1123))
	req.Header.Set("Accept", "application/vnd.api+json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		Logger.Printf("error fetching request")
		return err
	}
	defer func() {
		res.Body.Close()
	}()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return mapF3Error(res)
	}

	if body != nil {
		if err = json.NewDecoder(res.Body).Decode(body); err != nil {
			Logger.Printf("error decoding response body")
			return fmt.Errorf("error decoding json body. error: %w", err)
		}
	}

	return nil
}

func mapF3Error(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusBadRequest:
		var errRes = &F3StatusBadRequest{}
		if err := json.NewDecoder(res.Body).Decode(errRes); err == nil {
			return errRes
		} else {
			return fmt.Errorf("corrupted payload for bad request. %w", &F3StatusBadRequest{})
		}
	case http.StatusUnauthorized:
		return F3StatusUnauthorized
	case http.StatusForbidden:
		return F3StatusForbidden
	case http.StatusNotFound:
		return F3StatusNotFound
	case http.StatusMethodNotAllowed:
		return F3StatusMethodNotAllowed
	case http.StatusNotAcceptable:
		return F3StatusNotAcceptable
	case http.StatusConflict:
		return F3StatusConflict
	case http.StatusTooManyRequests:
		return F3StatusTooManyRequests
	case http.StatusInternalServerError:
		return F3StatusInternalServerError
	case http.StatusBadGateway:
		return F3StatusBadGateway
	case http.StatusServiceUnavailable:
		return F3StatusServiceUnavailable
	case http.StatusGatewayTimeout:
		return F3StatusGatewayTimeout
	default:
		return fmt.Errorf("status code: %d, unsupported error: %w",res.StatusCode, F3UnsupportedError)
	}
}

type PaginatedPayload struct {
	Data  []Data	`json:"data"`
	Links Links		`json:"links"`
}

type Payload struct {
	Data  Data  `json:"data"`
	Links Links `json:"links"`
}

type Data struct {
	Id             UUID       			`json:"id"`
	OrganisationId UUID       			`json:"organisation_id"`
	RecordType     string       		`json:"type"`
	Version        uint32       		`json:"version"`
	CreateOn       time.Time    		`json:"created_on"`
	ModifiedOn     time.Time    		`json:"modified_on"`
	Attributes     AccountAttributes 	`json:"attributes"`
}

type Links struct {
	Self string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type F3StatusBadRequest struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func (e *F3StatusBadRequest) Error() string {
	return fmt.Sprintf("Bad Request. Returned when the request submited doesn't meet server side validation"+
		"\nerrorCode: %q\nerrorMessage: %q", e.ErrorCode, e.ErrorMessage)
}

var F3StatusUnauthorized = fmt.Errorf("unauthorized. returned when trying to access api endpoints with an invalid or expired access token")
var F3StatusForbidden = fmt.Errorf("forbidden. returned when trying to obtain an access token with incorrect client credentials")
var F3StatusNotFound = fmt.Errorf("not found. returned when trying to access a non-existent endpoint or resource. returned in the validation api when a queried sort code cannot be found")
var F3StatusMethodNotAllowed = fmt.Errorf("method not allowed. returned when trying to access an endpoint that exists using a method that is not supported by the target resource")
var F3StatusNotAcceptable = fmt.Errorf("not acceptable. returned when trying to access content with an incorrect content type specific in the request header")
var F3StatusConflict = fmt.Errorf("conflict. the resource has already been created. it is safe ignore this error message and continue processing. returned for delete calls when an incorrect version has been specified")
var F3StatusTooManyRequests = fmt.Errorf("too many requests. returned when the rate limit for requests per second has been exceeded, please back-off immediately, then retry later")
var F3StatusInternalServerError = fmt.Errorf("server error. returned when an internal error occurs or the request times out. this is safe to retry after waiting a short amount of time")
var F3StatusBadGateway = fmt.Errorf("bad gateway. returned when there is a temporary internal networking problem. this is safe to retry after waiting a short amount of time")
var F3StatusServiceUnavailable = fmt.Errorf("service unavailable. returned when a service is temporarily overloaded. this is safe to retry after waiting a short amount of time")
var F3StatusGatewayTimeout = fmt.Errorf("gateway timeout. returned when there is a temporary internal networking problem. this is safe to retry after waiting a short amount of time")
var F3UnsupportedError = fmt.Errorf("the http status code returned is undocumented")
