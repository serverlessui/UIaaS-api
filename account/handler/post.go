package handler

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

//Account struct
type Account struct {
	AccountID string
	Name      string
}

//CreateAccount method to handle account creation requests
func CreateAccount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusCreated}, nil
}
