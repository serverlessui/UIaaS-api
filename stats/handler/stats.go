package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

//Visit is a struct representing visit statistical data
type Visit struct {
	Total           float64            `json:"total"`
	TotalPerMonth   []float64          `json:"totalPerMonth"`
	TotalLastWeek   float64            `json:"totalLastWeek"`
	TotalToday      float64            `json:"totalToday"`
	VisitsByCountry map[string]float64 `json:"visitsByCountry"`
}

//HandleGetStats method to handle get stats request
func HandleGetStats(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
