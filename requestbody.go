package gohttpclient

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

func (c *client) getRequestBody(body interface{}, contentType string) ([]byte, error) {

	switch assertedBody := body.(type) {
	case nil:
		return nil, nil
	case string:
		if assertedBody == "" {
			return nil, nil
		}
	case []byte:
		if len(assertedBody) == 0 {
			return nil, nil
		}
	}

	switch strings.ToLower(contentType) {

	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}
