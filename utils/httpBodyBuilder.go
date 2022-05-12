package utils

import (
	"bytes"
	"encoding/json"
)

func HttpBodyBuilder(arg any) *bytes.Buffer {
	bodyBytes, _ := json.Marshal(arg)
	body := bytes.NewBuffer(bodyBytes)

	return body
}
