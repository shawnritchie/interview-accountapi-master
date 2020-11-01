package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type createBuilder struct {
	AccountAttributes
	client *F3Client
	OrganisationId UUID
	Type string
	AccountId UUID
}

type CreateBuilder interface {
	WithCountry(country Country) CreateBuilder
	WithBaseCurrency(currency Currency) CreateBuilder
	WithBankId(bankId BankId) CreateBuilder
	WithBankIdCode(bankIdCode string) CreateBuilder
	WithAccountNumber(accountNumber string) CreateBuilder
	WithBic(bic SwiftCode) CreateBuilder
	WithIban(iban IBAN) CreateBuilder
	WithCustomerId(customerId string) CreateBuilder
	WithName(name Identifier) CreateBuilder
	WithAlternativeNames(alternativeName Identifier) CreateBuilder
	WithAccountClassification(classification Classification) CreateBuilder
	WithJointAccount(jointAccount bool) CreateBuilder
	WithAccountMatchingOptOut(optOut bool) CreateBuilder
	WithSecondaryIdentification(identifier Identifier) CreateBuilder
	WithSwitched(switched bool) CreateBuilder
	WithStatus(status Status) CreateBuilder
	WithOrganisationId(organisationId UUID) CreateBuilder
	WithAccountId(accountId UUID) CreateBuilder
	UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) CreateBuilder
	Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) CreateBuilder
	Validate(errors chan<- []error) CreateBuilder
}

func newAccountBuilder(client *F3Client) CreateBuilder {
	return createBuilder{
		client: client,
		Type: ACCOUNTTYPE,
	}
}

func (ab createBuilder) UnsafeRequest(ctx context.Context, response chan<- *Payload, errors chan<- []error) CreateBuilder {
	go ab.internalRequest(build(ab), ctx, response, errors)
	return ab
}

func (ab createBuilder) Validate(errors chan<- []error) CreateBuilder {
	if err := postValidators(ab); err != nil && len(err) > 0 {
		errors <- err
	}
	close(errors)
	return ab
}

func (ab createBuilder) Request(ctx context.Context, response chan<- *Payload, errors chan<- []error) CreateBuilder {
	if err := postValidators(ab); err != nil && len(err) > 0 {
		logPayloadErrors(err, response, errors)
	} else {
		go ab.internalRequest(build(ab), ctx, response, errors)
	}

	return ab
}

func (ab createBuilder) internalRequest(reqPayload * Payload, ctx context.Context, response chan<- *Payload, errors chan<- []error) {
	byteArray, err := json.Marshal(reqPayload)
	if err != nil {
		Logger.Println("error marshalling json payload for account creation")
		logPayloadError(fmt.Errorf("error marshalling payload: %+v - error: %w", reqPayload, err), response, errors)
		return
	}

	url := fmt.Sprintf("http://%s/v1/organisation/accounts", ab.client.Env.F3BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteArray))
	if err != nil {
		Logger.Printf("failed to creat new http request for %q", url)
		logPayloadError(fmt.Errorf("error creating request Method: 'Post' Url: %q - error: %w", url, err), response, errors)
		return
	}

	req = req.WithContext(ctx)
	res := &Payload{}
	if err := ab.client.request(req, res); err != nil {
		Logger.Printf("error requesting POST %q", url)
		logPayloadError(err, response, errors)
		return
	}

	logPayloadResponse(res, response, errors)
}



func (ab createBuilder) WithCountry(country Country) CreateBuilder {
	ab.Country = country
	return ab
}

func (ab createBuilder) WithBaseCurrency(currency Currency) CreateBuilder {
	ab.BaseCurrency = currency
	return ab
}

func (ab createBuilder) WithBankId(bankId BankId) CreateBuilder {
	ab.BankId = bankId
	return ab
}

func (ab createBuilder) WithBankIdCode(bankIdCode string) CreateBuilder {
	ab.BankIdCode = bankIdCode
	return ab
}

func (ab createBuilder) WithAccountNumber(accountNumber string) CreateBuilder {
	ab.AccountNumber = accountNumber
	return ab
}

func (ab createBuilder) WithBic(bic SwiftCode) CreateBuilder {
	ab.Bic = bic
	return ab
}

func (ab createBuilder) WithIban(iban IBAN) CreateBuilder {
	ab.Iban = iban
	return ab
}

func (ab createBuilder) WithCustomerId(customerId string) CreateBuilder {
	ab.CustomerId = customerId
	return ab
}

func (ab createBuilder) WithName(name Identifier) CreateBuilder {
	ab.Name = append(ab.Name, name)
	return ab
}

func (ab createBuilder) WithAlternativeNames(name Identifier) CreateBuilder {
	ab.AlternativeNames = append(ab.AlternativeNames, name)
	return ab
}

func (ab createBuilder) WithAccountClassification(classification Classification) CreateBuilder {
	ab.AccountClassification = classification
	return ab
}

func (ab createBuilder) WithJointAccount(jointAccount bool) CreateBuilder {
	ab.JointAccount = jointAccount
	return ab
}

func (ab createBuilder) WithAccountMatchingOptOut(optOut bool) CreateBuilder {
	ab.AccountMatchingOptOut = optOut
	return ab
}

func (ab createBuilder) WithSecondaryIdentification(identifier Identifier) CreateBuilder {
	ab.SecondaryIdentification = identifier
	return ab
}

func (ab createBuilder) WithSwitched(switched bool) CreateBuilder {
	ab.Switched = switched
	return ab
}

func (ab createBuilder) WithStatus(status Status) CreateBuilder {
	ab.Status = status
	return ab
}

func (ab createBuilder) WithOrganisationId(organisationId UUID) CreateBuilder {
	ab.OrganisationId = organisationId
	return ab
}

func (ab createBuilder) WithAccountId(accountId UUID) CreateBuilder {
	ab.AccountId = accountId
	return ab
}

