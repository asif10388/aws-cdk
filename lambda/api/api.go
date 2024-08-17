package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("Username or Password can't be empty")
	}

	userExists, err := api.dbStore.DoesUserExists(event.Username)

	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}

	if userExists {
		return fmt.Errorf("User already exists")
	}

	err = api.dbStore.RegisterUser(event)

	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}

	return nil
}
