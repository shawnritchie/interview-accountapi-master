package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type accountBuilder struct {
	AccountAttributes
	client *F3Client
	OrganisationId UUID
	Type string
	AccountId UUID
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
	UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) AccountBuilder
	Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) AccountBuilder
	Validate(errors chan<- []error) AccountBuilder
}

func newAccountBuilder(client *F3Client) AccountBuilder {
	return accountBuilder{
		client: client,
		Type: ACCOUNTTYPE,
	}
}

func (ab accountBuilder) UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) AccountBuilder {
	go ab.internalRequest(build(ab), ctx, response, errors)
	return ab
}

func (ab accountBuilder) Validate(errors chan<- []error) AccountBuilder {
	if err := postValidators(ab); err != nil && len(err) > 0 {
		errors <- err
	}
	close(errors)
	return ab
}

func (ab accountBuilder) Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) AccountBuilder {
	if err := postValidators(ab); err != nil && len(err) > 0 {
		logErrors(err, response, errors)
	} else {
		go ab.internalRequest(build(ab), ctx, response, errors)
	}

	return ab
}

func (ab accountBuilder) internalRequest(reqPayload * Payload, ctx context.Context, response chan<- *Payload, errors chan<- []error) {
	byteArray, err := json.Marshal(reqPayload)
	if err != nil {
		Logger.Println("error marshalling json payload for account creation")
		logError(fmt.Errorf("error marshalling payload: %+v - error: %w", reqPayload, err), response, errors)
		return
	}

	url := fmt.Sprintf("http://%s/v1/organisation/accounts", ab.client.Env.F3BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteArray))
	if err != nil {
		Logger.Printf("failed to creat new http request for %q", url)
		logError(fmt.Errorf("error creating request Method: 'Post' Url: %q - error: %w", url, err), response, errors)
		return
	}

	req = req.WithContext(ctx)
	att := &AccountAttributes{}
	res, err := ab.client.request(req, att)
	if err != nil {
		Logger.Printf("error requesting POST %q", url)
		logError(err, response, errors)
		return
	}

	logResponse(res, response, errors)
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

