package api

import "lambda-func/database"

type ApiHander struct {
	dbStore database.DynamoDBClient
}
