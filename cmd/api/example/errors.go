package example

import "errors"

var (
	FailedToRetrieveExampleData = errors.New("failed to retrieve example data from database")
	InvalidRequestPayload       = errors.New("invalid request payload provided")
)

const (
	InternalServerErrorMessage = "Failed to process the request"
	BadRequestMessage          = "Invalid request payload"
)
