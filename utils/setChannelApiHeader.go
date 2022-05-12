package utils

import (
	"net/http"
	"os"
)

func SetChannelApiHeader(req *http.Request) {
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"x-access-key": []string{os.Getenv("OPEN_API_ACCESS_KEY")},
		"x-access-secret": []string{os.Getenv("OPEN_API_ACCESS_SECRET")},
	}
}
