package form3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type fetchBuilder struct {
	AccountAttributes
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
	go fb.internalRequest(fb.AccountId, ctx, response, errors)
	return fb
}

func (fb fetchBuilder) Validate(errors chan<- []error) FetchBuilder {
	if err := fb.validate(); err != nil && len(err) > 0 {
		errors <- err
	}
	close(errors)
	return fb
}

func (fb fetchBuilder) Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) FetchBuilder {
	if err := fb.validate(); err != nil && len(err) > 0 {
		logErrors(err, response, errors)
		return fb
	} else {
		go fb.internalRequest(fb.AccountId, ctx, response, errors)
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

func (fb fetchBuilder) internalRequest(accountId UUID, ctx context.Context, response chan<- *Payload, errors chan<- []error) {
	if err := accountId.IsValid(); err != nil {
		Logger.Printf("UUID validation failed UUID: %s", accountId)
		logError(fmt.Errorf("invalid account id: %q. %w", accountId, err), response, errors)
		return
	}

	url := fmt.Sprintf("http://%s/v1/organisation/accounts/%s", fb.client.Env.F3BaseURL, url.QueryEscape(string(accountId)))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Logger.Printf("failed to creat new http request for %q", url)
		logError(fmt.Errorf("error creating request Method: 'GET' Url: %q - error: %w", url, err), response, errors)
		return
	}

	req = req.WithContext(ctx)
	att := &AccountAttributes{}
	res, err := fb.client.request(req, att)
	if err != nil {
		Logger.Printf("error requesting GET %q", url)
		logError(err, response, errors)
		return
	}

	logResponse(res, response, errors)
}