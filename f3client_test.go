package currency

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

type f3ClientState struct {
	F3Client       *F3Client
	organisationId string
	payload        *Payload
	errors          []error
}

func (state *f3ClientState) anInitiatedClient() error {
	f3Client, err := NewF3Client()
	if err != nil {
		return fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}
	state.F3Client = f3Client
	return nil
}

func (state *f3ClientState) anOrganisationIdOf(orgId string) error {
	state.organisationId = orgId
	return nil
}

func (state *f3ClientState) iSendRequestTo(method string, path string, accounts *messages.PickleStepArgument_PickleTable) error {

	builder := NewTypedAccountBuilder()
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
		}
	}

	att := builder.RawBuild()
	payload, err := state.F3Client.CreateAccount(context.Background(), uuid.New().String(), state.organisationId, &att)
	state.payload = payload
	if err != nil {
		state.errors = []error{err}
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
	if !containsErrorType(state.errors, f3StatusBadRequest) {
		return fmt.Errorf("was expecting 'F3StatusBadRequest' error but got %v", state.errors)
	}

	return nil
}

func containsErrorType(errs []error, target interface{}) bool {
	for _, err := range errs {
		if !errors.As(err, target) {
			return true
		}
	}
	return false
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	state := &f3ClientState{}

	ctx.BeforeScenario(func(sc *godog.Scenario) {
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


}
