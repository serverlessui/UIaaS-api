package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/serverlessui/UIaaS-api/stats/handler"
	"github.com/serverlessui/UIaaS-api/stats/repository"
	"github.com/serverlessui/UIaaS-api/stats/statistics"
)

const (
	tableName = "DYNAMODB_TABLE"
)

//CreateRepository is a method to instantiate an AccountRepository implementation
func CreateRepository() handler.StatsMetaDataRepository {
	tableName := os.Getenv(tableName)
	return repository.DynamoDBStatsMetaRepository{Svc: createDynamoDB(), TableName: tableName}
}

//CreateStatistics method to instantiate Statistics implementation
func CreateStatistics() handler.WebsiteStatistics {
	return statistics.AthenaStats{Client: createAthena()}
}

func createAthena() *athena.Athena {
	svc := athena.New(createSession())
	return svc
}
func createDynamoDB() *dynamodb.DynamoDB {

	return dynamodb.New(createSession())
}

func createSession() *session.Session {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(1)
	}
	return sess
}
