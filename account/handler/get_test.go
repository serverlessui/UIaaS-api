package handler

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetAccountStatusOK(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	resp, err := GetAccount(request)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := 200

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestGetAccountStatusAccountIDPassedInIsReturned(t *testing.T) {
	expected := "ACCOUNTID"
	pathParameterMap := make(map[string]string)
	pathParameterMap[accountIDParam] = expected
	request := events.APIGatewayProxyRequest{PathParameters: pathParameterMap}
	resp, err := GetAccount(request)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}
	var account Account
	err = json.Unmarshal([]byte(resp.Body), &account)
	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}
	actual := account.AccountID

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
