package handler

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

const (
	responseAccountID         = "ACCOUNTID"
	createAccountResponseBody = "{\"accountID\":\"ACCOUNTID\",\"name\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"companyName\":\"\",\"userName\":\"\"}"
)

func TestCreateAccountReturnsAccountPostHandlerReturns201(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mockAccountRepo := mockGoodAccountRepository{}

	resp, err := CreateAccount(request, mockAccountRepo)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.Body
	expected := createAccountResponseBody

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestCreateAccountReturnsAccountPostHandlerReturnsCorrectBody(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mockAccountRepo := mockGoodAccountRepository{}

	resp, err := CreateAccount(request, mockAccountRepo)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := http.StatusCreated

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestCreateAccountReturnsErrorPostHandlerReturns503(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mockAccountRepo := mockBadAccountRepository{}

	resp, err := CreateAccount(request, mockAccountRepo)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := http.StatusServiceUnavailable

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
