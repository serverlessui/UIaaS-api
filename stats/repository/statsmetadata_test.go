package repository

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

// Define a mock to return a basic success
type mockGoodDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (client mockGoodDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}

func (client mockGoodDynamoDbClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{}, nil
}

type mockBadDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (client mockBadDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, errors.New("Error")
}

func (client mockBadDynamoDbClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{}, errors.New("An error occured")
}
func TestWhenCreateMetaDataDynamoClientReturnsSuccesfullNoErrorIsReturned(t *testing.T) {
	repo := DynamoDBStatsMetaRepository{Svc: mockGoodDynamoDbClient{}, TableName: ""}
	datum := handler.StatsMetaData{}
	err := repo.CreateMetaData(&datum)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
}

func TestWhenCreateMetaDataDynamoClientReturnsErrorErrorIsReturned(t *testing.T) {
	repo := DynamoDBStatsMetaRepository{Svc: mockBadDynamoDbClient{}, TableName: ""}
	datum := handler.StatsMetaData{}
	err := repo.CreateMetaData(&datum)

	if err == nil {
		t.Log("Error not enountered when one was")
		t.Fail()
	}
}

func TestWhenGetMetaDataDynamoClientReturnsSuccesfullErrorIsNil(t *testing.T) {
	repo := DynamoDBStatsMetaRepository{Svc: mockGoodDynamoDbClient{}, TableName: ""}

	_, err := repo.GetMetaData("SITEID")

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
}

func TestWhenGetMetaDataDynamoClientReturnsErrorErrorIsReturned(t *testing.T) {
	repo := DynamoDBStatsMetaRepository{Svc: mockBadDynamoDbClient{}, TableName: ""}

	_, err := repo.GetMetaData("SITEID")

	if err == nil {
		t.Log("Error not encountered when one expected ")
		t.Fail()
	}
}
