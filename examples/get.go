package examples

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Get() (string, error) {
	resp, err := client.GET("https://api.github.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Status: ", resp.Status)
	/*
		var bodyJSON []byte

		err = json.Indent(bytes.NewBuffer(bodyJSON), bodyBytes, "", "	")
		if err != nil {
			return "", err
		}

		bodyString := string(bodyJSON)
		fmt.Println("Body: ", bodyString)
		return bodyString, nil
	*/

	bodyString := string(bodyBytes)
	fmt.Println("Body: ", bodyString)
	return bodyString, nil

}
