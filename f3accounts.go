package form3

type AccountAttributes struct {
	Country                 Country        `json:"country"`
	BaseCurrency            Currency       `json:"base_currency"`
	BankId                  BankId         `json:"bank_id"`
	BankIdCode              string         `json:"bank_id_code"`
	AccountNumber           string         `json:"account_number"`
	Bic                     SwiftCode      `json:"bic"`
	Iban                    IBAN           `json:"iban"`
	CustomerId              string         `json:"customer_id"`
	Name                    []Identifier   `json:"name"`
	AlternativeNames        []Identifier   `json:"alternative_names"`
	AccountClassification   Classification `json:"account_classification"`
	JointAccount            bool           `json:"joint_account"`
	AccountMatchingOptOut   bool           `json:"account_matching_opt_out"`
	SecondaryIdentification Identifier     `json:"secondary_identification"`
	Switched                bool           `json:"switched"`
	Status                  Status         `json:"status"`
}