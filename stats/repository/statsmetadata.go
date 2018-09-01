package repository

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

//DynamoDBStatsMetaRepository implementation of StatsMetaDataRepository
type DynamoDBStatsMetaRepository struct {
	Svc       dynamodbiface.DynamoDBAPI
	TableName string
}

//GetMetaData method to retrieve StatsMetaData for a site
func (repo DynamoDBStatsMetaRepository) GetMetaData(siteID string) (*handler.StatsMetaData, error) {
	metaData := handler.StatsMetaData{}

	result, err := repo.Svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"siteID": {
				S: aws.String(siteID),
			},
		},
	})

	if err != nil {
		fmt.Println("Got error calling Query:")
		fmt.Println(err.Error())
		return &metaData, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &metaData)

	if err != nil {
		fmt.Println("Error parsing response")
		fmt.Println(err.Error())
		return &metaData, err
	}
	return &metaData, nil
}

//CreateMetaData method to store StatsMetaData for a site
func (repo DynamoDBStatsMetaRepository) CreateMetaData(metadatum *handler.StatsMetaData) error {
	dynamoAccountEntity, err := dynamodbattribute.MarshalMap(metadatum)

	if err != nil {
		log.Println("Got error marshalling map:")
		log.Println(err.Error())
		return err
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
		return err
	}
	return nil
}
