package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

//Account struct
type Account struct {
	AccountID   string `json:"accountID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	CompanyName string `json:"companyName"`
	UserName    string `json:"userName"`
}

//AccountRepository interface defining methods to interact with Account entity
type AccountRepository interface {
	CreateAccount(account *Account) (*Account, error)
	GetAccount(accoundID string) (*Account, error)
}

//CreateAccount method to handle account creation requests
func CreateAccount(request events.APIGatewayProxyRequest, repository AccountRepository) (events.APIGatewayProxyResponse, error) {
	var account Account
	err := json.Unmarshal([]byte(request.Body), &account)
	returnedAccout, err := repository.CreateAccount(&account)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusServiceUnavailable}, nil
	}

	resp, err := MarshallAccountStructToString(returnedAccout)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: http.StatusCreated}, nil
}
