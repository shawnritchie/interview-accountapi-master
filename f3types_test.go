package currency

import (
	"testing"
)

func TestStatusValidation(t *testing.T) {
	statuses := []struct {
		scenario    string
		status      string
		expectError bool
	}{
		{"Valid Status", string(PENDING), false},
		{"Valid Status", string(CONFIRMED), false},
		{"Valid Status", string(FAILED), false},
		{"Invalid Status", "What", true},
		{"Empty BankId", "", true},
	}
	for _, s := range statuses {
		t.Run(s.scenario, func(t *testing.T) {
			status := Status(s.status)

			err := status.IsValid()
			switch s.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid Status %q", s.status)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid status %q", err, s.status)
				}
			}
		})
	}
}

func TestClassificationValidation(t *testing.T) {
	classifications := []struct {
		scenario       string
		classification string
		expectError    bool
	}{
		{"Valid Classification", string(PERSONAL), false},
		{"Valid Classification", string(BUSINESS), false},
		{"Invalid Status", "What", true},
		{"Empty BankId", "", true},
	}

	for _, c := range classifications {
		t.Run(c.scenario, func(t *testing.T) {
			classification := Classification(c.classification)

			err := classification.IsValid()
			switch c.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid Status %q", c.classification)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid classification %q", err, c.classification)
				}
			}
		})
	}
}

func TestBankIdValidation(t *testing.T) {
	bankIds := []struct {
		scenario    string
		bankId      string
		expectError bool
	}{
		{"Valid BankId", "valid", false},
		{"Valid BankId", "12345678901", false},
		{"Valid BankId", "1", false},
		{"Empty BankId", "", true},
		{"BankId too Long", "123456789012", true},
	}

	for _, bid := range bankIds {
		t.Run(bid.scenario, func(t *testing.T) {
			bankId := BankId(bid.bankId)

			err := bankId.IsValid()
			switch bid.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid Bank Id %q", bid.bankId)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid Bank Id %q", err, bid.bankId)
				}
			}
		})
	}
}

func TestSwiftCodeValidation(t *testing.T) {
	swiftCodes := []struct {
		scenario    string
		swiftCode   string
		expectError bool
	}{
		{"Valid SwiftCode", "12345678", false},
		{"Valid SwiftCode", "12345678901", false},
		{"Empty SwiftCode", "", true},
		{"SwiftCode too short", "1234567", true},
		{"SwiftCode too Long", "123456789012", true},
	}

	for _, sc := range swiftCodes {
		t.Run(sc.scenario, func(t *testing.T) {
			swiftCode := SwiftCode(sc.swiftCode)

			err := swiftCode.IsValid()
			switch sc.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid SwiftCode %q", sc.swiftCode)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid SwiftCode %q", err, sc.swiftCode)
				}
			}
		})
	}
}

func TestIBANValidation(t *testing.T) {
	ibans := []struct {
		scenario    string
		iban        string
		expectError bool
	}{
		{"Valid IBAN", "123456789012345678901234567890", false},
		{"Empty IBAN", "", true},
		{"IBAN too short", "1234567", true},
		{"IBAN too Long", "1234567890123456789012345678901234567890", true},
	}

	for _, i := range ibans {
		t.Run(i.scenario, func(t *testing.T) {
			iban := IBAN(i.iban)

			err := iban.IsValid()
			switch i.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid IBAN %q", i.iban)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid IBAN %q", err, i.iban)
				}
			}
		})
	}
}

func TestIdentifierValidation(t *testing.T) {
	largeString := func() (output string) {
		for len(output) <= 140 {
			output += "1"
		}
		return output
	}()

	identifiers := []struct {
		scenario    string
		identifier  string
		expectError bool
	}{
		{"Valid Identifier", "123456789012345678901234567890", false},
		{"Empty Identifier", "", true},
		{"Identifier too Long", largeString, true},
	}

	for _, i := range identifiers {
		t.Run(i.scenario, func(t *testing.T) {
			identifier := Identifier(i.identifier)

			err := identifier.IsValid()
			switch i.expectError {
			case true:
				if err == nil {
					t.Errorf("validation did not fail when it was expected too! invalid Identifier %q", i.identifier)
				}
			case false:
				if err != nil {
					t.Errorf("validation failed with %q on valid Identifier %q", err, i.identifier)
				}
			}
		})
	}
}
