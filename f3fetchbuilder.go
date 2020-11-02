package form3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type fetchBuilder struct {
	client *F3Client
	AccountId UUID
}

type FetchBuilder interface {
	WithAccountId(accountId UUID) FetchBuilder
	UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) FetchBuilder
	Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) FetchBuilder
	Validate(errors chan<- []error) FetchBuilder
}

func newFetchBuilder(client *F3Client) FetchBuilder {
	return fetchBuilder{
		client: client,
	}
}

func (fb fetchBuilder) WithAccountId(accountId UUID) FetchBuilder {
	fb.AccountId = accountId
	return fb
}

func (fb fetchBuilder) UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) FetchBuilder {
	go fb.internalRequest(ctx, response, errors)
	return fb
}

func (fb fetchBuilder) Validate(errors chan<- []error) FetchBuilder {
	if err := fb.validate(); len(err) > 0 {
		errors <- err
	}
	close(errors)
	return fb
}

func (fb fetchBuilder) Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) FetchBuilder {
	if err := fb.validate(); len(err) > 0 {
		logPayloadErrors(err, response, errors)
		return fb
	} else {
		go fb.internalRequest(ctx, response, errors)
		return fb
	}
}

func (fb fetchBuilder) validate() (errors []error) {
	if fb.client == nil {
		errors = append(errors, fmt.Errorf("F3Client not set"))
	}

	if fb.AccountId.IsZeroValue() {
		errors = append(errors, fmt.Errorf("missing account id in fetch request %w", accountIdFieldMissing))
	} else if err := fb.AccountId.IsValid(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (fb fetchBuilder) internalRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) {
	url := fmt.Sprintf("http://%s/v1/organisation/accounts/%s", fb.client.Env.F3BaseURL, url.QueryEscape(string(fb.AccountId)))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Logger.Printf("failed to creat new http request for %q", url)
		logPayloadError(fmt.Errorf("error creating request Method: 'GET' Url: %q - error: %w", url, err), response, errors)
		return
	}

	req = req.WithContext(ctx)
	res := &Payload{}
	if err := fb.client.request(req, res); err != nil {
		Logger.Printf("error requesting GET %q", url)
		logPayloadError(err, response, errors)
		return
	}

	logPayloadResponse(res, response, errors)
}