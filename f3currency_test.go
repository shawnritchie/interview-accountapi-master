package currency

import (
	"errors"
	"testing"
)

func TestCurrencyValidation(t *testing.T) {
	currencyCodes := []struct{
		scenario    string
		currencyStr string
		expectError bool
	}{
		{"Valid Currency", "EUR", false},
		{"Invalid Currency", "XYZ", true},
		{"Empty Currency", "", true},
	}
	var invalidCurrencyError *InvalidCurrency

	for _, cc := range currencyCodes {
		t.Run(cc.scenario, func(t *testing.T) {
			currency := NewCurrency(cc.currencyStr)

			err := currency.IsValid()
			switch cc.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid currency %q", cc.currencyStr)
				}

				if ok := errors.As(err, &invalidCurrencyError); !ok {
					t.Errorf("validation failed with incorrect error! invalid currency %q", cc.currencyStr)
				}
			case false:
				if err != nil {
					t.Errorf("Validation failed with %s on valid currency %q", err, cc.currencyStr)
				}
			}
		})
	}
}
