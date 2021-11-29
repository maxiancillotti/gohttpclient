package examples

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/maxiancillotti/gohttpclient/mock"
)

func TestGetExample(t *testing.T) {

	mock.MockupServer.Start()

	t.Run("TestGetExampleFetchingFromAPI", func(t *testing.T) {

		// Initialization
		mock.MockupServer.DeleteMocks()
		mock.MockupServer.AddMock(mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  fmt.Errorf("Timeout fetching from API"),
		})

		// Execution
		resp, err := GetExample()

		// Validation
		if resp != "" {
			t.Error("empty response value was expected")
		}
		if err == nil {
			t.Error("an error was expected")
		}
		if err.Error() != "Timeout fetching from API" {
			t.Error("error message received is invalid:", err.Error())
		}
	})
	/*
		t.Run("TestGetExample", func(t *testing.T) {
			// Execution
			resp, err := GetExample()

			// Validation
			if err != nil {
				t.Error("Error executing Get:", err)
			}
			fmt.Println("Get response from testing:", resp)
		})

		t.Run("", func(t *testing.T) {
			// Execution
			resp, err := GetExample()

			// Validation
			if err != nil {
				t.Error("Error executing Get:", err)
			}
			fmt.Println("Get response from testing:", resp)
		})
	*/
}
