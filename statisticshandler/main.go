package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Visit is a struct representing visit statistical data
type Visit struct {
	Total           float64            `json:"total"`
	TotalPerMonth   []float64          `json:"totalPerMonth"`
	TotalLastWeek   float64            `json:"totalLastWeek"`
	TotalToday      float64            `json:"totalToday"`
	VisitsByCountry map[string]float64 `json:"visitsByCountry"`
}

//GetStatistics is the entry point to the Lambda function
func GetStatistics(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	visit := &Visit{
		Total:         10000000,
		TotalLastWeek: 10000,
		TotalToday:    3333,
	}
	resp, err := json.Marshal(visit)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	responseString := string(resp)

	return events.APIGatewayProxyResponse{Body: responseString, StatusCode: 200}, nil
}

func main() {
	lambda.Start(GetStatistics)
}
