package handler

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

const (
	bucketNameParam = "bucketname"
)

//WebsiteStatistics interface defining CRUD ops arounds stats
type WebsiteStatistics interface {
	Get(sourceName string) (*VisitStatistics, error)
}

//VisitStatistics is a struct representing visit statistical data
type VisitStatistics struct {
	Total           int                  `json:"total"`
	TotalPerMonth   []SiteVisitsPerMonth `json:"totalPerMonth"`
	TotalLastWeek   float64              `json:"totalLastWeek"`
	TotalToday      float64              `json:"totalToday"`
	VisitsByCountry map[string]float64   `json:"visitsByCountry"`
}

//SiteVisitsPerMonth represents month visits per site
type SiteVisitsPerMonth struct {
	Data  []int  `json:"data"`
	Label string `json:"label"`
}

//HandleGetStats method to handle get stats request
func HandleGetStats(request *events.APIGatewayProxyRequest, stats WebsiteStatistics) (events.APIGatewayProxyResponse, error) {
	site1StatPerMonth := SiteVisitsPerMonth{Data: []int{0, 5000, 15000, 8000, 15000, 9000, 30000}, Label: "CCB Technologies"}

	bucket := request.QueryStringParameters[bucketNameParam]
	log.Println("received request with bucket ", bucket)
	visit, err := stats.Get(bucket)

	if err != nil {
		log.Println("ERROR: error received when processing request ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 503, Headers: createCorsHeaders()}, nil
	}
	visit.TotalPerMonth = []SiteVisitsPerMonth{site1StatPerMonth}

	resp, err := json.Marshal(visit)

	if err != nil {
		log.Println("ERROR: error marshalling response ", err)

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
