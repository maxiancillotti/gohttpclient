package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	// Initialization

	// Execution
	resp, err := Get()

	// Validation
	if err != nil {
		t.Error("Error executing Get:", err)
	}
	fmt.Println("Get response from testing:", resp)
}
