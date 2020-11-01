package form3

import (
	"fmt"
	"net/http"
)

const (
	CREATE = http.MethodPost
	PATCH  = http.MethodPatch
	ACCOUNTTYPE = "accounts"
)

func NewAccountBuilder() AccountBuilder {
	return accountBuilder{Type: ACCOUNTTYPE}
}

type accountBuilder struct {
	AccountAttributes
	OrganisationId UUID
	Type string
	AccountId UUID
}

func (ab accountBuilder) RawBuild() *Payload {
	return build(ab)
}

func (ab accountBuilder) Build(requestType string) (*Payload, []error) {
	switch requestType {
	case CREATE:
		if err := postValidators(ab); err != nil && len(err) > 0 {
			return nil, err
		}
	default:
		return nil, []error{ fmt.Errorf("unsupported requested type %q. supported types %v", requestType, []string{CREATE}) }
	}

	return build(ab), nil
}

func build(u accountBuilder) *Payload {
	return &Payload{
		Data: Data{
			Id:             u.AccountId,
			OrganisationId: u.OrganisationId,
			RecordType:     u.Type,
			Attributes:     AccountAttributes{
				Country:                 u.Country,
				BaseCurrency:            u.BaseCurrency,
				BankId:                  u.BankId,
				BankIdCode:              u.BankIdCode,
				AccountNumber:           u.AccountNumber,
				Bic:                     u.Bic,
				Iban:                    u.Iban,
				CustomerId:              u.CustomerId,
				Name:                    u.Name,
				AlternativeNames:        u.AlternativeNames,
				AccountClassification:   u.AccountClassification,
				JointAccount:            u.JointAccount,
				AccountMatchingOptOut:   u.AccountMatchingOptOut,
				SecondaryIdentification: u.SecondaryIdentification,
				Switched:                u.Switched,
				Status:                  u.Status,
			},
		},
	}
}

type AccountBuilder interface {
	WithCountry(country Country) AccountBuilder
	WithBaseCurrency(currency Currency) AccountBuilder
	WithBankId(bankId BankId) AccountBuilder
	WithBankIdCode(bankIdCode string) AccountBuilder
	WithAccountNumber(accountNumber string) AccountBuilder
	WithBic(bic SwiftCode) AccountBuilder
	WithIban(iban IBAN) AccountBuilder
	WithCustomerId(customerId string) AccountBuilder
	WithName(name Identifier) AccountBuilder
	WithAlternativeNames(alternativeName Identifier) AccountBuilder
	WithAccountClassification(classification Classification) AccountBuilder
	WithJointAccount(jointAccount bool) AccountBuilder
	WithAccountMatchingOptOut(optOut bool) AccountBuilder
	WithSecondaryIdentification(identifier Identifier) AccountBuilder
	WithSwitched(switched bool) AccountBuilder
	WithStatus(status Status) AccountBuilder
	WithOrganisationId(organisationId UUID) AccountBuilder
	WithAccountId(accountId UUID) AccountBuilder
	RawBuild() *Payload
	Build(requestType string) (*Payload, []error)
}

func (ab accountBuilder) WithCountry(country Country) AccountBuilder {
	ab.Country = country
	return ab
}

func (ab accountBuilder) WithBaseCurrency(currency Currency) AccountBuilder {
	ab.BaseCurrency = currency
	return ab
}

func (ab accountBuilder) WithBankId(bankId BankId) AccountBuilder {
	ab.BankId = bankId
	return ab
}

func (ab accountBuilder) WithBankIdCode(bankIdCode string) AccountBuilder {
	ab.BankIdCode = bankIdCode
	return ab
}

func (ab accountBuilder) WithAccountNumber(accountNumber string) AccountBuilder {
	ab.AccountNumber = accountNumber
	return ab
}

func (ab accountBuilder) WithBic(bic SwiftCode) AccountBuilder {
	ab.Bic = bic
	return ab
}

func (ab accountBuilder) WithIban(iban IBAN) AccountBuilder {
	ab.Iban = iban
	return ab
}

func (ab accountBuilder) WithCustomerId(customerId string) AccountBuilder {
	ab.CustomerId = customerId
	return ab
}

func (ab accountBuilder) WithName(name Identifier) AccountBuilder {
	ab.Name = append(ab.Name, name)
	return ab
}

func (ab accountBuilder) WithAlternativeNames(name Identifier) AccountBuilder {
	ab.AlternativeNames = append(ab.AlternativeNames, name)
	return ab
}

func (ab accountBuilder) WithAccountClassification(classification Classification) AccountBuilder {
	ab.AccountClassification = classification
	return ab
}

func (ab accountBuilder) WithJointAccount(jointAccount bool) AccountBuilder {
	ab.JointAccount = jointAccount
	return ab
}

func (ab accountBuilder) WithAccountMatchingOptOut(optOut bool) AccountBuilder {
	ab.AccountMatchingOptOut = optOut
	return ab
}

func (ab accountBuilder) WithSecondaryIdentification(identifier Identifier) AccountBuilder {
	ab.SecondaryIdentification = identifier
	return ab
}

func (ab accountBuilder) WithSwitched(switched bool) AccountBuilder {
	ab.Switched = switched
	return ab
}

func (ab accountBuilder) WithStatus(status Status) AccountBuilder {
	ab.Status = status
	return ab
}

func (ab accountBuilder) WithOrganisationId(organisationId UUID) AccountBuilder {
	ab.OrganisationId = organisationId
	return ab
}
func (ab accountBuilder) WithAccountId(accountId UUID) AccountBuilder {
	ab.AccountId = accountId
	return ab
}


