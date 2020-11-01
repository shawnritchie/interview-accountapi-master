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
	os.Setenv("F3BaseURL", "localhost:8080")
	os.Setenv("F3Timeout", "60s")
	os.Setenv("F3MaxRetries", "3")
}

func validAccount(f3Client *F3Client) AccountBuilder {
	return f3Client.Create().
		WithCountry(Countries["GB"]).
		WithBankId("000006").
		WithBic("NWBKGB22").
		WithBankIdCode("GBDSC").
		WithAccountClassification("Personal")
}

type f3ClientState struct {
	F3Client       *F3Client
	organisationId UUID
	accountId      UUID
	payload        *Payload
	errors         []error
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
	if att, ok := state.payload.Data.Attributes.(*AccountAttributes); !ok {
		return fmt.Errorf("error extracting attributes from payload")
	} else {
		if att.AccountNumber == "" {
			return fmt.Errorf("expecting account numer to be generated from %q", state.F3Client.Env.F3BaseURL)
		}
	}
	return nil
}

func (state *f3ClientState) responseShouldContainAnIBANNumber() error {
	if att, ok := state.payload.Data.Attributes.(*AccountAttributes); !ok {
		return fmt.Errorf("error extracting attributes from payload")
	} else {
		if att.Iban.IsZeroValue() {
			return fmt.Errorf("expecting iban to be generated from %q", state.F3Client.Env.F3BaseURL)
		}
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

func InitializeScenario(ctx *godog.ScenarioContext) {
	state := &f3ClientState{}

	ctx.BeforeScenario(func(sc *godog.Scenario) {
		state.organisationId = ""
		state.accountId = ""
		state.payload = nil
		state.errors = nil
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
}
