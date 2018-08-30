package repository

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/serverlessui/UIaaS-api/account/handler"
)

const (
	resultAccountID = "ACCOUNTID"
)

// Define a mock to return a basic success
type mockGoodDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}
type mockBadDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (client mockGoodDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}
func (client mockBadDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, errors.New("An error occured")
}

func (client mockGoodDynamoDbClient) Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return &dynamodb.QueryOutput{}, nil
}
func (client mockBadDynamoDbClient) Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return &dynamodb.QueryOutput{}, errors.New("An error occured")
}

func TestGetAccountReturnsAccountWithSuccessfulQueryCall(t *testing.T) {
	repo := DynamoAccountRepository{Svc: mockGoodDynamoDbClient{}}
	accountID := "accountid"

	actual, err := repo.GetAccount(accountID)
	if err != nil {
		t.Log("Error received when none expected ", err)
	}

	if actual == nil {
		t.Log("response was nil when not expected to be")
		t.Fail()
	}
}

func TestGetAccountReturnsErrorWhenQueryCallFails(t *testing.T) {
	repo := DynamoAccountRepository{Svc: mockBadDynamoDbClient{}}
	accountID := "accountid"

	_, err := repo.GetAccount(accountID)
	if err == nil {
		t.Log("Error received when none expected ", err)
	}

}
func TestSourceDynamoDBRepositoryCreateAccountReturnsUUID(t *testing.T) {
	repo := DynamoAccountRepository{Svc: mockGoodDynamoDbClient{}}
	account := handler.Account{}
	response, err := repo.CreateAccount(&account)
	actual := response.AccountID
	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}
	if actual == "" {
		t.Log("accountID not set correctly")
		t.Fail()
	}
}

func TestSourceDynamoDBRepositoryCreateAccountErrorsReturnsError(t *testing.T) {
	repo := DynamoAccountRepository{Svc: mockBadDynamoDbClient{}}
	account := handler.Account{}
	_, err := repo.CreateAccount(&account)

	if err == nil {
		t.Log("No error was returned")
		t.Fail()
	}
}
