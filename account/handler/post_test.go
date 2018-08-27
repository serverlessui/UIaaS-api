package handler

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCreateAccount(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	resp, err := CreateAccount(request)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := 201

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
