package form3

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/google/uuid"
	"os"
)

func init() {
	_ = os.Setenv("F3BaseURL", "localhost:8080")
	_ = os.Setenv("F3Timeout", "60s")
	_ = os.Setenv("F3MaxRetries", "3")
}

func validAccount(f3Client *F3Client) CreateBuilder {
	return f3Client.Create().
		WithCountry(Countries["GB"]).
		WithBankId("000006").
		WithBic("NWBKGB22").
		WithBankIdCode("GBDSC").
		WithAccountClassification("Personal")
}

type f3ClientState struct {
	F3Client         *F3Client
	Paginator        Paginator
	organisationId   UUID
	accountId        UUID
	payload          *Payload
	PaginatedPayload *PaginatedPayload
	errors           []error
}

func (state *f3ClientState) anInitiatedClient() error {
	f3Client, err := NewF3Client()
	if err != nil {
		return fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}
	state.F3Client = f3Client
	return nil
}

func containsErrorType(errs []error, target interface{}) bool {
	for _, err := range errs {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

func containsError(errs []error, target error) bool {
	for _, err := range errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (state *f3ClientState) anOrganisationIdOf(orgId string) error {
	state.organisationId = UUID(orgId)
	return nil
}

func (state *f3ClientState) iSendRequestTo(method string, path string, accounts *messages.PickleStepArgument_PickleTable) error {
	builder := state.F3Client.Create().
		WithAccountId(state.accountId).
		WithOrganisationId(state.organisationId)

	for i := 1; i < len(accounts.Rows); i++ {
		key := accounts.Rows[i].Cells[0].Value
		value := accounts.Rows[i].Cells[1].Value

		switch key {
		case "Country":
			builder = builder.WithCountry(Country(value))
		case "BankId":
			builder = builder.WithBankId(BankId(value))
		case "BIC":
			builder = builder.WithBic(SwiftCode(value))
		case "BankIdCode":
			builder = builder.WithBankIdCode(value)
		case "AccountNumber":
			builder = builder.WithAccountNumber(value)
		case "IBAN":
			builder = builder.WithIban(IBAN(value))
		case "Classification":
			builder = builder.WithAccountClassification(Classification(value))
		case "Name":
			builder = builder.WithName(Identifier(value))
		case "AlternativeNames":
			builder = builder.WithAlternativeNames(Identifier(value))
		case "SecondaryIdentification":
			builder = builder.WithSecondaryIdentification(Identifier(value))
		case "Status":
			builder = builder.WithStatus(Status(value))
		}
	}

	response := make(chan *Payload, 1)
	errors := make(chan []error, 1)

	builder.UnsafeRequest(context.Background(), response, errors)

	err := <-errors
	state.payload = <-response

	if err != nil {
		state.errors = err
	}

	return nil
}

func (state *f3ClientState) responseShouldContainAnAccountNumber() error {
	if state.payload.Data.Attributes.AccountNumber == "" {
		return fmt.Errorf("expecting account numer to be generated from %q", state.F3Client.Env.F3BaseURL)
	}
	return nil
}

func (state *f3ClientState) responseShouldContainAnIBANNumber() error {
	if state.payload.Data.Attributes.Iban.IsZeroValue() {
		return fmt.Errorf("expecting iban to be generated from %q", state.F3Client.Env.F3BaseURL)
	}
	return nil
}

func (state *f3ClientState) weExpectAValidResponse() error {
	if state.errors != nil && len(state.errors) > 0 {
		return fmt.Errorf("errors: %v", state.errors)
	}

	return nil
}

func (state *f3ClientState) weExpectABadRequestResponse() error {
	var f3StatusBadRequest *F3StatusBadRequest
	if !containsErrorType(state.errors, &f3StatusBadRequest) {
		return fmt.Errorf("was expecting 'F3StatusBadRequest' error but got %v", state.errors)
	}

	return nil
}

func (state *f3ClientState) aRandomOrganisationId() error {
	state.organisationId = UUID(uuid.New().String())
	return nil
}

func (state *f3ClientState) aRandomAccountId() error {
	state.accountId = UUID(uuid.New().String())
	return nil
}

func (state *f3ClientState) aValidAccountHasBeenRegistered() error {
	response := make(chan *Payload, 1)
	errors := make(chan []error, 1)

	validAccount(state.F3Client).
		WithOrganisationId(state.organisationId).
		WithAccountId(state.accountId).
		UnsafeRequest(context.Background(), response, errors)

	err := <-errors
	state.payload = <-response

	if err != nil {
		return fmt.Errorf("multiple errors encountered during request %+v", err)
	}

	return nil
}

func (state *f3ClientState) anInvalidOrganisationId() error {
	state.organisationId = "123456"
	return nil
}

func (state *f3ClientState) anInvalidAccountId() error {
	state.accountId = "123456"
	return nil
}

func (state *f3ClientState) iSendRequestToWithTheAccountId(arg1, arg2 string) error {
	response := make(chan *Payload, 1)
	errors := make(chan []error, 1)

	state.F3Client.Fetch().
		WithAccountId(state.accountId).
		UnsafeRequest(context.Background(), response, errors)

	state.errors = <-errors
	state.payload = <-response

	return nil
}

func (state *f3ClientState) weExpectAValidationError() error {
	if state.errors != nil && len(state.errors) > 0 {
		return fmt.Errorf("errors: %v", state.errors)
	}

	return nil
}

func (state *f3ClientState) weExpectAHttpStatusCodeNotFound() error {
	if !containsError(state.errors, F3StatusNotFound) {
		return fmt.Errorf("we were expecting 404 not found error. yet we failed with %v", state.errors)
	}

	return nil
}

func (state *f3ClientState) iSendRequestToForPageWithAPageSizeOf(method, path string, page, size int) error {
	response := make(chan *PaginatedPayload, 1)
	errors := make(chan []error, 1)

	state.Paginator = state.F3Client.
							List().
							WithPage(page).
							WithPageSize(size).
							Request(context.Background(), response, errors)

	state.errors = <-errors
	state.PaginatedPayload = <-response

	return nil
}

func (state *f3ClientState) iNavigateToTheNextPage() error {
	response := make(chan *PaginatedPayload, 1)
	errors := make(chan []error, 1)

	state.Paginator = state.Paginator.Next(context.Background(), response, errors)

	state.errors = <-errors
	state.PaginatedPayload = <-response

	return nil
}

func (state *f3ClientState) iNavigateToThePreviousPage() error {
	response := make(chan *PaginatedPayload, 1)
	errors := make(chan []error, 1)

	state.Paginator = state.Paginator.Prev(context.Background(), response, errors)

	state.errors = <-errors
	state.PaginatedPayload = <-response

	return nil
}

func (state *f3ClientState) iNavigateToTheFirstPage() error {
	response := make(chan *PaginatedPayload, 1)
	errors := make(chan []error, 1)

	state.Paginator = state.Paginator.First(context.Background(), response, errors)

	state.errors = <-errors
	state.PaginatedPayload = <-response

	return nil
}

func (state *f3ClientState) iNavigateToTheLastPage() error {
	response := make(chan *PaginatedPayload, 1)
	errors := make(chan []error, 1)

	state.Paginator = state.Paginator.Last(context.Background(), response, errors)

	state.errors = <-errors
	state.PaginatedPayload = <-response

	return nil
}

func (state *f3ClientState) aPageSizeOf(pageSize int) error {
	if state.PaginatedPayload == nil {
		return fmt.Errorf("no list response found we were expecting a response containing a list of accounts")
	}

	if len(state.PaginatedPayload.Data) != pageSize {
		return fmt.Errorf("list response contained a total of %d accounts. we expected %d accounts with the specified page size",
			len(state.PaginatedPayload.Data), pageSize)
	}

	return nil
}

func (state *f3ClientState) iSendRequestToWithTheAccountIdAndVersion(method, path string, version int) error {
	errors := make(chan []error, 1)

	state.F3Client.Delete().
		WithAccountId(state.accountId).
		WithVersion(version).
		UnsafeRequest(context.Background(), errors)

	state.errors = <-errors

	return nil
}

func (state *f3ClientState) weExpectAHttpStatusCodeConflict() error {
	if !containsError(state.errors, F3StatusConflict) {
		return fmt.Errorf("we were expecting 409 conflict error. yet we failed with %v", state.errors)
	}

	return nil
}

func (state *f3ClientState) weExpectNoValidationErrors() error  {
	if state.errors != nil && len(state.errors) > 0 {
		return fmt.Errorf("we were expecting no validation errors yet we encountered %v", state.errors)
	}
	return nil
}

func (state *f3ClientState) weExpectValidationErrors() error {
	if state.errors == nil && len(state.errors) == 0 {
		return fmt.Errorf("we were expecting validation errors yet we didn't encountered any")
	}
	return nil
}

func (state *f3ClientState) weValidateTheAccountBuilderWithProperties(requestType string, accounts *messages.PickleStepArgument_PickleTable) error {
	builder := state.F3Client.Create().
		WithAccountId(state.accountId).
		WithOrganisationId(state.organisationId)

	for i := 1; i < len(accounts.Rows); i++ {
		key := accounts.Rows[i].Cells[0].Value
		value := accounts.Rows[i].Cells[1].Value

		switch key {
		case "Country":
			builder = builder.WithCountry(Country(value))
		case "BankId":
			builder = builder.WithBankId(BankId(value))
		case "BIC":
			builder = builder.WithBic(SwiftCode(value))
		case "BankIdCode":
			builder = builder.WithBankIdCode(value)
		case "AccountNumber":
			builder = builder.WithAccountNumber(value)
		case "IBAN":
			builder = builder.WithIban(IBAN(value))
		case "Classification":
			builder = builder.WithAccountClassification(Classification(value))
		case "Name":
			builder = builder.WithName(Identifier(value))
		case "AlternativeNames":
			builder = builder.WithAlternativeNames(Identifier(value))
		case "SecondaryIdentification":
			builder = builder.WithSecondaryIdentification(Identifier(value))
		case "Status":
			builder = builder.WithStatus(Status(value))
		}
	}

	errors := make(chan []error, 1)
	builder.Validate(errors)
	state.errors = <-errors

	return nil
}


func InitializeScenario(ctx *godog.ScenarioContext) {
	state := &f3ClientState{}

	ctx.BeforeScenario(func(sc *godog.Scenario) {
		state.organisationId = ""
		state.accountId = ""
		state.payload = nil
		state.errors = nil
		state.Paginator = nil
		state.PaginatedPayload = nil
	})

	ctx.Step(`^an initiated Client$`, state.anInitiatedClient)
	ctx.Step(`^an organisationId of "([^"]*)"$`, state.anOrganisationIdOf)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, state.iSendRequestTo)
	ctx.Step(`^response should contain an account Number$`, state.responseShouldContainAnAccountNumber)
	ctx.Step(`^response should contain an IBAN Number$`, state.responseShouldContainAnIBANNumber)
	ctx.Step(`^we expect a valid response$`, state.weExpectAValidResponse)
	ctx.Step(`^we expect a bad request response$`, state.weExpectABadRequestResponse)
	ctx.Step(`^we expect no validation errors$`, state.weExpectNoValidationErrors)
	ctx.Step(`^we expect validation errors$`, state.weExpectValidationErrors)
	ctx.Step(`^we validate the "([^"]*)" account builder with properties$`, state.weValidateTheAccountBuilderWithProperties)
	ctx.Step(`^a random organisation id$`, state.aRandomOrganisationId)
	ctx.Step(`^a valid account has been registered$`, state.aValidAccountHasBeenRegistered)
	ctx.Step(`^an invalid organisation id$`, state.anInvalidOrganisationId)
	ctx.Step(`^an invalid accountId$`, state.anInvalidAccountId)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with the accountId$`, state.iSendRequestToWithTheAccountId)
	ctx.Step(`^we expect a validation error$`, state.weExpectAValidationError)
	ctx.Step(`^we expect a http status code: Not Found$`, state.weExpectAHttpStatusCodeNotFound)
	ctx.Step(`^a random accountId$`, state.aRandomAccountId)
	ctx.Step(`^a random organisationId$`, state.aRandomOrganisationId)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" for page (\d+) with a page size of (\d+)$`, state.iSendRequestToForPageWithAPageSizeOf)
	ctx.Step(`^I navigate to the next page$`, state.iNavigateToTheNextPage)
	ctx.Step(`^I navigate to the previous page$`, state.iNavigateToThePreviousPage)
	ctx.Step(`^I navigate to the first page$`, state.iNavigateToTheFirstPage)
	ctx.Step(`^I navigate to the last page$`, state.iNavigateToTheLastPage)
	ctx.Step(`^a page size of (\d+)$`, state.aPageSizeOf)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with the accountId and version (\d+)$`, state.iSendRequestToWithTheAccountIdAndVersion)
	ctx.Step(`^we expect a http status code: Conflict$`, state.weExpectAHttpStatusCodeConflict)
}
