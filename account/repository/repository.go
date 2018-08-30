package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	uuid "github.com/satori/go.uuid"
	"github.com/serverlessui/UIaaS-api/account/handler"
)

//DynamoAccountRepository implementation of AccountRepository interface
type DynamoAccountRepository struct {
	Svc       dynamodbiface.DynamoDBAPI
	TableName string
}

//CreateAccount method to create account object
func (repo DynamoAccountRepository) CreateAccount(account *handler.Account) (*handler.Account, error) {
	u1 := uuid.Must(uuid.NewV4(), nil)
	account.AccountID = u1.String()
	dynamoAccountEntity, err := dynamodbattribute.MarshalMap(account)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
	}

	// Create item in table
	input := &dynamodb.PutItemInput{
		Item:      dynamoAccountEntity,
		TableName: aws.String(repo.TableName),
	}
	//put item
	_, err = repo.Svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return &handler.Account{}, err
	}

	return account, nil
}

//GetAccount returns an Account object based on an ID
func (repo DynamoAccountRepository) GetAccount(accoundID string) (*handler.Account, error) {
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(repo.TableName),
		// IndexName: aws.String("accountID"),
		KeyConditions: map[string]*dynamodb.Condition{
			"accountID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(accoundID),
					},
				},
			},
		},
	}
	var resp, err = repo.Svc.Query(queryInput)

	account := handler.Account{}

	if err != nil {
		fmt.Println("Got error calling Query:")
		fmt.Println(err.Error())
		return &account, err
	}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &account)

	if err != nil {
		fmt.Println("Error parsing response")
		fmt.Println(err.Error())
		return &account, err
	}
	return &account, nil
}
