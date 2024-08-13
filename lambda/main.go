package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string
}

func HandleRequest(event Event) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("Username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by %s", event.Username), nil

}

func main() {
	lambda.Start(HandleRequest)
}
