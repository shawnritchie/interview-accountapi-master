package form3

import (
	"fmt"
	"regexp"
)

const (
	PENDING   Status = "pending"
	CONFIRMED Status = "confirmed"
	FAILED    Status = "failed"

	PERSONAL Classification = "Personal"
	BUSINESS Classification = "Business"
)

type TypeValidator interface {
	IsValid() error
	IsZeroValue() bool
}

type Classification string

var zeroValueClassification = Classification("")

func (c *Classification) IsValid() error {
	switch *c {
	case PERSONAL, BUSINESS:
		return nil
	}

	return fmt.Errorf("invalid classification %q. only acceptable values are %v", *c, []Classification{PERSONAL, BUSINESS})
}

func (c *Classification) IsZeroValue() bool {
	return zeroValueClassification == *c
}

type Status string

var zeroValueStatus = Status("")

func (s *Status) IsValid() error {
	switch *s {
	case PENDING, CONFIRMED, FAILED:
		return nil
	}
	return fmt.Errorf("invalid status %q. only acceptable values are %v", *s, []Status{PENDING, CONFIRMED, FAILED})
}

func (s *Status) IsZeroValue() bool {
	return zeroValueStatus == *s
}

type SwiftCode string

var swiftCodeRegex = "^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$"
var zeroValueSwiftCode = SwiftCode("")

func (sc *SwiftCode) IsValid() error {
	match, _ := regexp.MatchString(swiftCodeRegex, string(*sc))
	if !match {
		return fmt.Errorf("invalid swift code %q, swiftcode should match %q", *sc, swiftCodeRegex)
	}

	return nil
}

func (sc *SwiftCode) IsZeroValue() bool {
	return zeroValueSwiftCode == *sc
}

type IBAN string

var ibanRegex string = "^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$"
var zeroValueIBAN = IBAN("")

func (i *IBAN) IsValid() error {
	if l := len(*i); l < 16 || l > 34 {
		return fmt.Errorf("invalid IBAN %q. Length: %d - min minLength 16 characters", *i, len(*i))
	}

	match, _ := regexp.MatchString(ibanRegex, string(*i))
	if !match {
		return fmt.Errorf("invalid iban %q, iban should match %q", *i, ibanRegex)
	}

	return nil
}

func (i *IBAN) IsZeroValue() bool {
	return zeroValueIBAN == *i
}

type Identifier string

var zeroValueIdentifier = Identifier("")

func (i *Identifier) IsValid() error {
	if l := len(*i); l == 0 || l > 140 {
		return fmt.Errorf("invalid Identifier %q. Length: %d - max minLength 140 characters", *i, len(*i))
	}
	return nil
}

func (i *Identifier) IsZeroValue() bool {
	return zeroValueIdentifier == *i
}

type BankId string

var zeroValueBankId = BankId("")

func (bid *BankId) IsValid() error {
	if l := len(*bid); l == 0 || l > 11 {
		return fmt.Errorf("invalid Bank Id %q. Length: %d - Max Length 11 characters", *bid, len(*bid))
	}
	return nil
}

func (bid *BankId) IsZeroValue() bool {
	return zeroValueBankId == *bid
}

type UUID string

var uuidValidation = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
var zeroValueUUID = UUID("")

func (u *UUID) IsValid() error {
	match, _ := regexp.MatchString(uuidValidation, string(*u))
	if !match {
		return fmt.Errorf("uuid: %q did not match regex expression %q", *u, uuidValidation)
	}

	return nil
}

func (u *UUID) IsZeroValue() bool {
	return zeroValueUUID == *u
}
