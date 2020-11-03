package form3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type deleteBuilder struct {
	client    *F3Client
	AccountId UUID
	Version   int
}

type DeleteBuilder interface {
	WithAccountId(accountId UUID) DeleteBuilder
	WithVersion(version int) DeleteBuilder
	UnsafeRequest(ctx context.Context, errors chan<- []error) DeleteBuilder
	Request(ctx context.Context, errors chan<- []error) DeleteBuilder
	Validate(errors chan<- []error) DeleteBuilder
}

func newDeleteBuilder(client *F3Client) DeleteBuilder {
	return deleteBuilder{
		client: client,
	}
}

func (d deleteBuilder) WithAccountId(accountId UUID) DeleteBuilder {
	d.AccountId = accountId
	return d
}

func (d deleteBuilder) WithVersion(version int) DeleteBuilder {
	d.Version = version
	return d
}

func (d deleteBuilder) UnsafeRequest(ctx context.Context, errors chan<- []error) DeleteBuilder {
	go d.internalRequest(ctx, errors)
	return d
}

func (d deleteBuilder) Request(ctx context.Context, errors chan<- []error) DeleteBuilder {
	if err := d.validate(); len(err) > 0 {
		logErrors(err, errors)
		return d
	} else {
		go d.internalRequest(ctx, errors)
		return d
	}
}

func (d deleteBuilder) Validate(errors chan<- []error) DeleteBuilder {
	if err := d.validate(); len(err) > 0 {
		errors <- err
	}
	close(errors)
	return d
}

func (d deleteBuilder) validate() (errors []error) {
	if d.client == nil {
		errors = append(errors, fmt.Errorf("F3Client not set"))
	}

	if d.AccountId.IsZeroValue() {
		errors = append(errors, fmt.Errorf("missing account id in fetch request %w", accountIdFieldMissing))
	} else if err := d.AccountId.IsValid(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (d deleteBuilder) internalRequest(ctx context.Context, errors chan<- []error) {
	url := fmt.Sprintf("http://%s/v1/organisation/accounts/%s?version=%d",
		d.client.Env.F3BaseURL, url.QueryEscape(string(d.AccountId)), d.Version)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		Logger.Printf("failed to creat new delete request for %q", url)
		logError(fmt.Errorf("error posting request Method: 'DELETE' Url: %q - error: %w", url, err), errors)
		return
	}

	req = req.WithContext(ctx)
	if err := d.client.request(req, nil); err != nil {
		Logger.Printf("error requesting GET %q", url)
		logError(err, errors)
		return
	}

	close(errors)
}

func logError(err error, errors chan<- []error) {
	errors <- []error{err}
	close(errors)
}

func logErrors(errs []error, errors chan<- []error) {
	errors <- errs
	close(errors)
}
