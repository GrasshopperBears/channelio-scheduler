package structs

type RequestType uint8

const (
	REQUEST_ERROR RequestType = 0
	REQUEST_HELP RequestType = 1
	REQUEST_ADD RequestType = 2
	REQUEST_GET RequestType = 3
	REQUEST_DELETE RequestType = 4
)
