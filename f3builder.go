package form3

const (
	ACCOUNTS = "accounts"
)

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