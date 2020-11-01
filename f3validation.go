package form3

import (
	"errors"
	"fmt"
	"regexp"
)

type Validator func(createBuilder) []error
type BuilderConstraint Validator

var accountIdFieldMissing = missingFieldError("id")
var countryFieldMissing = missingFieldError("country")
var bicFieldMissing = missingFieldError("bic")
var bankIdFieldMissing = missingFieldError("bank_id")
var classificationFieldMissing = missingFieldError("account_classification")
var TooManyNames = errors.New("names array is restricted to a maximum string[4]")
var TooManyAlternativeNames = errors.New("alternative names array is restricted to a maximum string[3]")

func missingFieldError(field string) error {
	return fmt.Errorf("%q field required inside of CreateBuilder", field)
}

func postValidators(ab createBuilder) []error {
	switch ab.Country {
	case "GB":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bicValidator,
			bankIdCodeValidator, accountNumberValidator)(ab)
	case "AU":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdCodeValidator, bicValidator,
			accountNumberValidator, emptyIbanValidator)(ab)
	case "BE":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "CA":
		return composeValidators(validateSetFields, validateRequiredFields, bicValidator, bankIdCodeValidator,
			accountNumberValidator, emptyIbanValidator)(ab)
	case "FR":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "DE":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "GR":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "HK":
		return composeValidators(validateSetFields, validateRequiredFields, bicValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "IT":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator, italyValidator)(ab)
	case "LU":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "NL":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bicValidator,
			bankIdCodeValidator, accountNumberValidator)(ab)
	case "PL":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "PT":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "ES":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "CH":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bankIdCodeValidator,
			accountNumberValidator)(ab)
	case "US":
		return composeValidators(validateSetFields, validateRequiredFields, bankIdValidator, bicValidator,
			bankIdCodeValidator, accountNumberValidator, emptyIbanValidator)(ab)
	default:
		return composeValidators(validateSetFields, validateRequiredFields)(ab)
	}
}

func emptyIbanValidator(ab createBuilder) (errors []error) {
	if ab.Iban != "" {
		errors = append(errors, fmt.Errorf("iban should be empty"))
	}
	return errors
}

func composeValidators(validators ...Validator) Validator {
	f := func(ab createBuilder) (errors []error) {
		for _, fx := range validators {
			if err := fx(ab); err != nil {
				errors = append(errors, err...)
			}
		}
		return errors
	}
	return f
}

func validateSetFields(ab createBuilder) (errors []error) {
	if err := ab.Country.IsValid(); !ab.Country.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.BaseCurrency.IsValid(); !ab.BaseCurrency.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.BankId.IsValid(); !ab.BankId.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.Bic.IsValid(); !ab.Bic.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.Iban.IsValid(); !ab.Iban.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.AccountClassification.IsValid(); !ab.AccountClassification.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.SecondaryIdentification.IsValid(); !ab.SecondaryIdentification.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if err := ab.Status.IsValid(); !ab.Status.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	if l := len(ab.Name); l > 0 {
		errors = append(errors, nameValidator(ab)...)
	}

	if l := len(ab.AlternativeNames); l > 0 {
		errors = append(errors, alternativeNameValidator(ab)...)
	}

	return errors
}

func nameValidator(ab createBuilder) (errors []error) {
	if l := len(ab.Name); l > 4 {
		errors = append(errors, TooManyNames)
	}

	for _, id := range ab.Name {
		if err := id.IsValid(); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func alternativeNameValidator(ab createBuilder) (errors []error) {
	if l := len(ab.AlternativeNames); l > 3 {
		errors = append(errors, TooManyNames)
	}

	for _, id := range ab.AlternativeNames {
		if err := id.IsValid(); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func bicValidator(ab createBuilder) (errors []error) {
	if ab.Bic.IsZeroValue() {
		errors = append(errors, bicFieldMissing)
	}

	if err := ab.Bic.IsValid(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func bankIdValidator(ab createBuilder) (errors []error) {
	if validator, ok := bankIdValidationMap[ab.Country]; ok {
		for _, err := range stringValidator(string(ab.BankId), validator) {
			errors = append(errors, fmt.Errorf("invalid 'BankId', error: %w", err))
		}
	} else if err := ab.BankId.IsValid(); !ab.BankId.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	return errors
}

func accountNumberValidator(ab createBuilder) (errors []error) {
	if validator, ok := accountNumberLengthMap[ab.Country]; ok {
		for _, err := range stringValidator(string(ab.AccountNumber), validator) {
			errors = append(errors, fmt.Errorf("invalid 'Account Number', error: %w", err))
		}
	} else if err := ab.BankId.IsValid(); !ab.BankId.IsZeroValue() && err != nil {
		errors = append(errors, err)
	}

	return errors
}

func stringValidator(data string, validator stringValidation) (errors []error) {
	if validator.shouldBeEmpty() && data != "" {
		errors = append(errors, fmt.Errorf("field should be empty"))
	} else if data != "" && (len(data) < validator.minLength || len(data) > validator.maxLength) {
		errors = append(errors, fmt.Errorf("field validation failed. string length requirements min: %d max: %d",
			validator.minLength, validator.maxLength))
	}

	if validator.required && data == "" {
		errors = append(errors, fmt.Errorf("field should not be empty"))
	}

	if validator.regex != "" && data != "" {
		match, _ := regexp.MatchString(validator.regex, string(data))
		if !match {
			errors = append(errors, fmt.Errorf("field did not match regex expression %q", validator.regex))
		}
	}

	return errors
}

func bankIdCodeValidator(ab createBuilder) (errors []error) {
	if expectedCodes, ok := BankIdCodes[ab.Country]; ok {
		if ab.BankIdCode != expectedCodes {
			errors = append(errors, fmt.Errorf("invalid bank id code: %q for country %q should be %q", ab.BankIdCode, ab.Country, expectedCodes))
		}
	}
	return errors
}

func italyValidator(ab createBuilder) (errors []error) {
	if ab.AccountNumber == "" && len(ab.BankId) != 10 {
		errors = append(errors, fmt.Errorf("invalid Italian Bank Id %q. seeing no account number is submited length should be 10 characters", ab.BankId))
	}

	if ab.AccountNumber != "" && len(ab.BankId) != 11 {
		errors = append(errors, fmt.Errorf("invalid Italian Bank Id %q. seeing an account number is submited length should be 11 characters", ab.BankId))
	}

	return errors
}

func validateRequiredFields(ab createBuilder) (errors []error) {
	if ab.Country.IsZeroValue() {
		errors = append(errors, countryFieldMissing)
	} else if err := ab.Country.IsValid(); err != nil {
		errors = append(errors, err)
	}

	if ab.AccountClassification.IsZeroValue() {
		errors = append(errors, classificationFieldMissing)
	}

	if err := ab.AccountClassification.IsValid(); err != nil {
		errors = append(errors, err)
	}

	if err := ab.OrganisationId.IsValid(); err != nil {
		errors = append(errors, err)
	}

	if err := ab.AccountId.IsValid(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

type stringValidation struct {
	required  bool
	minLength int
	maxLength int
	regex     string
}

func (sv *stringValidation) shouldBeEmpty() bool {
	return sv.minLength == -1 && sv.maxLength == -1
}

func exactLengthValidation(required bool, length int) stringValidation {
	return stringValidation{
		required:  required,
		minLength: length,
		maxLength: length,
	}
}

func regexValidation(required bool, length int, regex string) stringValidation {
	return stringValidation{
		required:  required,
		minLength: length,
		maxLength: length,
		regex:     regex,
	}
}

var emptyBankId = stringValidation{
	required:  false,
	minLength: -1,
	maxLength: -1,
}

var BankIdCodes = map[Country]string{
	Countries["GB"]: "GBDSC",
	Countries["AU"]: "AUBSB",
	Countries["BE"]: "BE",
	Countries["CA"]: "CACPA",
	Countries["FR"]: "FR",
	Countries["DE"]: "DEBLZ",
	Countries["GR"]: "GRBIC",
	Countries["HK"]: "HKNCC",
	Countries["IT"]: "ITNCC",
	Countries["LU"]: "LULUX",
	Countries["NL"]: "",
	Countries["PL"]: "PLKNR",
	Countries["PT"]: "PTNCC",
	Countries["ES"]: "ESNCC",
	Countries["CH"]: "CHBCC",
	Countries["US"]: "USABA",
}

var bankIdValidationMap = map[Country]stringValidation{
	Countries["GB"]: regexValidation(true, 6, "^[0-9]{6}$"),
	Countries["AU"]: regexValidation(false, 6, "^[0-9]{6}$"),
	Countries["BE"]: exactLengthValidation(true, 3),
	Countries["CA"]: regexValidation(false, 9, "^0.{8}$"),
	Countries["FR"]: exactLengthValidation(true, 10),
	Countries["DE"]: exactLengthValidation(true, 8),
	Countries["GR"]: exactLengthValidation(true, 7),
	Countries["HK"]: exactLengthValidation(false, 3),
	Countries["IT"]: stringValidation{
		required:  true,
		minLength: 10,
		maxLength: 11,
	},
	Countries["LU"]: exactLengthValidation(true, 3),
	Countries["NL"]: emptyBankId,
	Countries["PL"]: exactLengthValidation(true, 8),
	Countries["PT"]: exactLengthValidation(true, 8),
	Countries["ES"]: exactLengthValidation(true, 8),
	Countries["CH"]: exactLengthValidation(true, 5),
	Countries["US"]: exactLengthValidation(true, 9),
}

var accountNumberLengthMap = map[Country]stringValidation{
	Countries["GB"]: exactLengthValidation(false, 8),
	Countries["AU"]: stringValidation{
		required:  false,
		minLength: 6,
		maxLength: 10,
		regex:     "^(?!0).{6,10}$",
	},
	Countries["BE"]: exactLengthValidation(false, 7),
	Countries["CA"]: stringValidation{
		required:  false,
		minLength: 7,
		maxLength: 12,
	},
	Countries["FR"]: exactLengthValidation(false, 10),
	Countries["DE"]: exactLengthValidation(false, 7),
	Countries["GR"]: exactLengthValidation(false, 16),
	Countries["HK"]: stringValidation{
		required:  false,
		minLength: 9,
		maxLength: 12,
	},
	Countries["IT"]: exactLengthValidation(false, 12),
	Countries["LU"]: exactLengthValidation(false, 13),
	Countries["NL"]: exactLengthValidation(false, 10),
	Countries["PL"]: exactLengthValidation(false, 16),
	Countries["PT"]: exactLengthValidation(false, 11),
	Countries["ES"]: exactLengthValidation(false, 10),
	Countries["CH"]: exactLengthValidation(false, 12),
	Countries["US"]: stringValidation{
		required:  false,
		minLength: 6,
		maxLength: 17,
	},
}
