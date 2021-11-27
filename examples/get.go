package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetExample() (string, error) {
	resp, err := httpClient.GET("https://api.github.com", nil)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Status: ", resp.Status)

	var bodyJSONBuf bytes.Buffer
	err = json.Indent(&bodyJSONBuf, bodyBytes, "", "\t")
	if err != nil {
		return "", err
	}

	fmt.Println("Body: ", bodyJSONBuf.String())
	return bodyJSONBuf.String(), nil
}
