package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shawnritchie/interview-accountapi-master"
	"os"
	"time"
)

func init() {
	if os.Getenv("F3BaseURL") == "" {
		_ = os.Setenv("F3BaseURL", "localhost:8080")
	}
}

func main() {
	f3Client, err := form3.NewF3Client()
	if err != nil {
		panic(fmt.Errorf("F3Client could not be initalized please check environmental variables 'F3BaseURL' is set up correct %w", err))
	}

	builder := f3Client.Create().
		WithCountry(form3.Countries["GB"]).
		WithBankId("000006").
		WithBic("NWBKGB22").
		WithBankIdCode("GBDSC").
		WithAccountClassification("Personal")

	for i := 0; i < 1500; i++ {
		response := make(chan *form3.Payload, 1)
		errors := make(chan []error, 1)

		builder.
			WithOrganisationId(form3.UUID(uuid.New().String())).
			WithAccountId(form3.UUID(uuid.New().String())).
			UnsafeRequest(context.Background(), response, errors)

		time.Sleep(10 * time.Millisecond)
	}
}
