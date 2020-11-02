package form3

import (
	"context"
	"fmt"
)


func ExampleF3CreateAccount() {
	f3Client, err := NewF3Client()
	if err != nil {
		fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}

	response := make(chan *Payload, 1)
	errors := make(chan []error, 1)

	f3Client.Create().
		WithAccountId("81d62ace-23f2-4aff-a7d6-60d7674bc5bb").
		WithOrganisationId("ea68b98a-471a-4c71-ac83-0f96a2bee973").
		WithCountry(Countries["GB"]).
		WithBankId("000006").
		WithBic("NWBKGB22").
		WithBankIdCode("GBDSC").
		WithAccountClassification("Personal").
		Request(context.Background(), response, errors)

	<- response
	<- errors
}

func ExampleF3FetchAccount() {
	f3Client, err := NewF3Client()
	if err != nil {
		fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}

	response := make(chan *Payload, 1)
	errors := make(chan []error, 1)

	f3Client.Fetch().
		WithAccountId("81d62ace-23f2-4aff-a7d6-60d7674bc5bb").
		Request(context.Background(), response, errors)

	<- response
	<- errors
}


func ExampleF3ListAccount() {
	f3Client, err := NewF3Client()
	if err != nil {
		fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}

	firstResponse := make(chan *PaginatedPayload, 1)
	firstErrors := make(chan []error, 1)

	lastResponse := make(chan *PaginatedPayload, 1)
	lastErrors := make(chan []error, 1)

	f3Client.List().
		WithPageSize(10).
		WithPage(0).
		Request(context.Background(), firstResponse, firstErrors).
		Last(context.Background(), lastResponse, lastErrors)

	<- firstResponse
	<- firstErrors
	<- lastResponse
	<- lastErrors
}

func ExampleF3DeleteAccount() {
	f3Client, err := NewF3Client()
	if err != nil {
		fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err)
	}

	errors := make(chan []error, 1)

	f3Client.Delete().
		WithAccountId("81d62ace-23f2-4aff-a7d6-60d7674bc5bb").
		WithVersion(0).
		Request(context.Background(), errors)

	<- errors
}