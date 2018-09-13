package handler

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

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

//StatisticsRun is a struct
type StatisticsRun struct {
	ExecutionID string
}

//StatisticsGenerator is an interface to handle creation of statistics
type StatisticsGenerator interface {
	Get(serviceID string) (*StatisticsRun, error)
}

//HandleGetStats method to handle get stats request
func HandleGetStats(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	site1StatPerMonth := SiteVisitsPerMonth{Data: []float64{0, 5000, 15000, 8000, 15000, 9000, 30000}, Label: "CCB Technologies"}

	visit := &VisitStatistics{
		Total:         10000000,
		TotalLastWeek: 10000,
		TotalToday:    3333,
		TotalPerMonth: []SiteVisitsPerMonth{site1StatPerMonth},
	}
	resp, err := json.Marshal(visit)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500, Headers: createCorsHeaders()}, nil
	}
	responseString := string(resp)

	return events.APIGatewayProxyResponse{Body: responseString, StatusCode: 200, Headers: createCorsHeaders()}, nil
}

//HandlePostStatistics method to handle creation of Stats request
func HandlePostStatistics(request events.SNSEvent, statsGenerator StatisticsGenerator) {
	log.Println("SNS reuest received")
	log.Println(request)
}

func createCorsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
}
