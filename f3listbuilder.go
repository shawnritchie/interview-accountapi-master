package form3

import (
	"context"
	"fmt"
	"net/http"
)

type listBuilder struct {
	AccountAttributes
	client   *F3Client
	Page     int
	PageSize int
	response *PaginatedPayload
}

type ListBuilder interface {
	WithPage(page int) ListBuilder
	WithPageSize(pageSize int) ListBuilder
	Validate(errors chan<- []error) ListBuilder
	UnsafeRequest(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
	Request(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
}

type Paginator interface {
	Next(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
	Prev(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
	First(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
	Last(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator
}

func newListBuilder(client *F3Client) ListBuilder {
	return listBuilder{
		client:   client,
		Page:     0,
		PageSize: 100,
	}
}

func (l listBuilder) WithPage(page int) ListBuilder {
	l.Page = page
	return l
}

func (l listBuilder) WithPageSize(pageSize int) ListBuilder {
	l.PageSize = pageSize
	return l
}

func (l listBuilder) Validate(errors chan<- []error) ListBuilder {
	if err := l.validate(); len(err) > 0 {
		errors <- err
		close(errors)
		return l
	}

	return l
}

func (l listBuilder) UnsafeRequest(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	url := fmt.Sprintf("http://%s/v1/organisation/accounts?page%%5Bnumber%%5D=%d&page%%5Bsize%%5D]=%d", l.client.Env.F3BaseURL, l.Page, l.PageSize)
	return l.internalRequest(url, ctx, response, errors)
}

func (l listBuilder) Request(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	if err := l.validate(); len(err) > 0 {
		errors <- err
		close(errors)
		return l
	} else {
		url := fmt.Sprintf("http://%s/v1/organisation/accounts?page%%5Bnumber%%5D=%d&page%%5Bsize%%5D]=%d", l.client.Env.F3BaseURL, l.Page, l.PageSize)
		Logger.Printf("URL: %s", url)
		return l.internalRequest(url, ctx, response, errors)
	}
}

func (l listBuilder) Next(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	if err := l.canPaginate(); err != nil {
		logPaginatedError(err, response, errors)
		return l
	}

	if l.response.Links.Next == "" {
		logPaginatedError(fmt.Errorf("response is missing next link, which is required for traversal"), response, errors)
		return l
	}

	url := fmt.Sprintf("http://%s%s", l.client.Env.F3BaseURL, l.response.Links.Next)
	l.internalRequest(url, ctx, response, errors)
	return l
}

func (l listBuilder) Prev(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	if err := l.canPaginate(); err != nil {
		logPaginatedError(err, response, errors)
		return l
	}

	if l.response.Links.Prev == "" {
		logPaginatedError(fmt.Errorf("response is missing next previous, which is required for traversal"), response, errors)
		return l
	}

	url := fmt.Sprintf("http://%s%s", l.client.Env.F3BaseURL, l.response.Links.Prev)
	l.internalRequest(url, ctx, response, errors)
	return l
}

func (l listBuilder) First(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	if err := l.canPaginate(); err != nil {
		logPaginatedError(err, response, errors)
		return l
	}

	if l.response.Links.First == "" {
		logPaginatedError(fmt.Errorf("response is missing next first, which is required for traversal"), response, errors)
		return l
	}

	url := fmt.Sprintf("http://%s%s", l.client.Env.F3BaseURL, l.response.Links.First)
	l.internalRequest(url, ctx, response, errors)
	return l
}

func (l listBuilder) Last(ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) Paginator {
	if err := l.canPaginate(); err != nil {
		logPaginatedError(err, response, errors)
		return l
	}

	if l.response.Links.Last == "" {
		logPaginatedError(fmt.Errorf("response is missing last link, which is required for traversal"), response, errors)
		return l
	}

	url := fmt.Sprintf("http://%s%s", l.client.Env.F3BaseURL, l.response.Links.Last)
	l.internalRequest(url, ctx, response, errors)
	return l
}

func (l listBuilder) canPaginate() error {
	if l.response == nil {
		return fmt.Errorf("list builder has no context. list builder needs to get context via Request or unsafeRequest")
	}
	return nil
}

func (l listBuilder) validate() (errors []error) {
	if l.client == nil {
		errors = append(errors, fmt.Errorf("F3Client not set"))
	}

	if l.Page < 0 {
		errors = append(errors, fmt.Errorf("page requested cannot be smaller then 0"))
	}

	if l.PageSize < 1 {
		errors = append(errors, fmt.Errorf("page size cannot be smaller then 1"))
	}

	return errors
}

func (l listBuilder) internalRequest(url string, ctx context.Context, response chan<- *PaginatedPayload, errors chan<- []error) listBuilder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Logger.Printf("failed to fetch request for %q", url)
		logPaginatedError(fmt.Errorf("error creating request Method: 'GET' Url: %q - error: %w", url, err), response, errors)
		return l
	}

	req = req.WithContext(ctx)
	res := &PaginatedPayload{}
	if err := l.client.request(req, res); err != nil {
		Logger.Printf("error requesting GET %q", url)
		logPaginatedError(err, response, errors)
		return l
	}
	l.response = res

	logPaginatedResponse(res, response, errors)
	return l
}

func logPaginatedError(err error, response chan<- *PaginatedPayload, errors chan<- []error) {
	close(response)
	errors <- []error{err}
	close(errors)
}

func logPaginatedResponse(payload *PaginatedPayload, response chan<- *PaginatedPayload, errors chan<- []error) {
	close(errors)
	response <- payload
	close(response)
}
