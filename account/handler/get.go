package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const (
	accountIDParam = "aid"
)

//GetAccount method to handle account retrieval requests
func GetAccount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accountID := request.PathParameters[accountIDParam]

	account := Account{AccountID: accountID, Name: "account name"}

	resp, err := json.Marshal(account)
	if err != nil {
		fmt.Println("Error processing request: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: http.StatusOK}, nil
}
