package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/athena"
)

const (
	bucketNameParam = "bucketname"
)

//WebsiteStatistics interface defining CRUD ops arounds stats
type WebsiteStatistics interface {
	Get(sourceName string) (*athena.ResultSet, error)
}

//VisitStatistics is a struct representing visit statistical data
type VisitStatistics struct {
	Total           float64              `json:"total"`
	TotalPerMonth   []SiteVisitsPerMonth `json:"totalPerMonth"`
	TotalLastWeek   float64              `json:"totalLastWeek"`
	TotalToday      float64              `json:"totalToday"`
	VisitsByCountry map[string]float64   `json:"visitsByCountry"`
}

//SiteVisitsPerMonth represents month visits per site
type SiteVisitsPerMonth struct {
	Data  []float64 `json:"data"`
	Label string    `json:"label"`
}

//HandleGetStats method to handle get stats request
func HandleGetStats(request *events.APIGatewayProxyRequest, stats WebsiteStatistics) (events.APIGatewayProxyResponse, error) {
	bucket := request.QueryStringParameters[bucketNameParam]
	visit, err := stats.Get(bucket)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 503, Headers: createCorsHeaders()}, nil
	}
	resp, err := json.Marshal(visit)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500, Headers: createCorsHeaders()}, nil
	}
	responseString := string(resp)

	return events.APIGatewayProxyResponse{Body: responseString, StatusCode: 200, Headers: createCorsHeaders()}, nil
}

func createCorsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
}
