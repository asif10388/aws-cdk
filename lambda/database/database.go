package database

type DynamoDBClient struct {
	databaseStore string
}

func NewDynamoDBClient() DynamoDBClient {
	return DynamoDBClient{
		databaseStore: "DBStore",
	}
}
