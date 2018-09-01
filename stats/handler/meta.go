package handler

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

const (
	siteIDParam = "sid"
)

//StatsMetaDataRepository interface for interacting with stats meta data entity
type StatsMetaDataRepository interface {
	CreateMetaData(metadatum *StatsMetaData) error
	GetMetaData(siteID string) (*StatsMetaData, error)
}

//StatsMetaData struct representing meta data of stats for site
type StatsMetaData struct {
	AccountID        string  `json:"accountID"`
	SiteID           string  `json:"siteID"`
	TableName        string  `json:"tableName"`
	CreatedTimestamp float64 `json:"createdTimestamp"`
}

//HandleGetStatsMeta is a method to handle Get stats request
func HandleGetStatsMeta(request *events.APIGatewayProxyRequest, repo StatsMetaDataRepository) (events.APIGatewayProxyResponse, error) {
	siteID := request.PathParameters[siteIDParam]

	response, err := repo.GetMetaData(siteID)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 503}, nil

	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error marshalling response ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
}

//HandlePostStatsMeta is a method to handle Post stats request
func HandlePostStatsMeta(request *events.APIGatewayProxyRequest, repo StatsMetaDataRepository) (events.APIGatewayProxyResponse, error) {
	var metaDatum StatsMetaData

	err := json.Unmarshal([]byte(request.Body), &metaDatum)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil

	}
	err = repo.CreateMetaData(&metaDatum)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 503}, nil

	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 201}, nil
}
