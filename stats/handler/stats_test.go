package handler

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleGetStats(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(request)
	if err != nil {
		t.Log("Error encounterd when none expected ", err)
		t.Fail()
	}

	expected := "{\"total\":10000000,\"totalPerMonth\":null,\"totalLastWeek\":10000,\"totalToday\":3333,\"visitsByCountry\":null}"
	actual := res.Body

	if actual != expected {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}

func TestHandleGetStatsHasCorsHeaders(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(request)
	if err != nil {
		t.Log("Error encounterd when none expected ", err)
		t.Fail()
	}

	expected := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
	actual := res.Headers

	if !reflect.DeepEqual(actual, expected) {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}
