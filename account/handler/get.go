package handler

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const (
	accountIDParam = "aid"
)

//GetAccount method to handle account retrieval requests
func GetAccount(request events.APIGatewayProxyRequest, repository AccountRepository) (events.APIGatewayProxyResponse, error) {
	accountID := request.PathParameters[accountIDParam]

	account, err := repository.GetAccount(accountID)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusServiceUnavailable}, nil
	}

	resp, err := MarshallAccountStructToString(account)
	if err != nil {
		fmt.Println("Error processing request: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: resp, StatusCode: http.StatusOK}, nil
}
