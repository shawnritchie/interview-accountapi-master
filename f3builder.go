package form3

const (
	ACCOUNTTYPE = "accounts"
)

func build(u createBuilder) *Payload {
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

func logPayloadError(err error, response chan<- *Payload, errors chan<- []error) {
	close(response)
	errors <- []error { err }
	close(errors)
}

func logPayloadErrors(err []error, response chan<- *Payload, errors chan<- []error) {
	close(response)
	errors <- err
	close(errors)
}

func logPayloadResponse(payload *Payload, response chan<- *Payload, errors chan<- []error) {
	close(errors)
	response <- payload
	close(response)
}