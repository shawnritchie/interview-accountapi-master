package form3

import (
	"fmt"
	"github.com/cucumber/messages-go/v10"
)

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
	builder := NewAccountBuilder()
	builder = builder.WithAccountId(state.accountId)
	builder = builder.WithOrganisationId(state.organisationId)

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

	_, err := builder.Build(requestType)
	state.errors = err

	return nil
}