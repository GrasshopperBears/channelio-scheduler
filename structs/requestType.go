package structs

type RequestType uint8

const (
	REQUEST_ERROR RequestType = 0
	REQUEST_ADD RequestType = 1
	REQUEST_GET RequestType = 2
	REQUEST_DELETE RequestType = 3
)
