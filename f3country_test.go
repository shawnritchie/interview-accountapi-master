package form3

import (
	"errors"
	"testing"
)

func TestZeroValueCurrency(t *testing.T) {
	if c, _ := Countries["XX"]; c != zeroValueCountry {
		t.Errorf("non existant country should return zero value")
	}
}

func TestCountryValidation(t *testing.T) {
	countryCodes := []struct{
		scenario    string
		countryStr  string
		expectError bool
	}{
		{"Valid Country", "FR", false},
		{"Invalid Country", "XX", true},
		{"Empty Country", "", true},
	}
	var invalidCountryError *InvalidCountry

	for _, cc := range countryCodes {
		t.Run(cc.scenario, func(t *testing.T) {
			country := NewCountry(cc.countryStr)

			err := country.IsValid()
			switch cc.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid local %q", cc.countryStr)
				}

				if ok := errors.As(err, &invalidCountryError); !ok {
					t.Errorf("validation failed with incorrect error! invalid local %q", cc.countryStr)
				}
			case false:
				if err != nil {
					t.Errorf("Validation failed with %s on valid local %q", err, cc.countryStr)
				}
			}
		})
	}
}
