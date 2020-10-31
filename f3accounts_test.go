package currency

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

	_, err := builder.Build(requestType)
	state.errors = err

	return nil
}


func containsError(errors []error, fx func(error) bool) bool {
	for _, e := range errors {
		if fx(e) {
			return true
		}
	}
	return false
}