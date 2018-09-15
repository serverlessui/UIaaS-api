package handler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/athena"
)

type mockWebsitStats struct {
}

func (mock mockWebsitStats) Get(sourceName string) (*athena.ResultSet, error) {
	return &athena.ResultSet{}, nil
}

type mockBadWebsitStats struct {
}

func (mock mockBadWebsitStats) Get(sourceName string) (*athena.ResultSet, error) {
	return &athena.ResultSet{}, errors.New("")
}
func TestHandleGetStats(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(&request, mockWebsitStats{})
	if err != nil {
		t.Log("Error encounterd when none expected ", err)
		t.Fail()
	}

	expected := "{\"ResultSetMetadata\":null,\"Rows\":null}"
	actual := res.Body

	if actual != expected {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestHandleGetStatsReturns503WhenGetStatsFails(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(&request, mockBadWebsitStats{})
	if err != nil {
		t.Log("Error encounterd when none expected ", err)
		t.Fail()
	}

	expected := 503
	actual := res.StatusCode

	if actual != expected {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestHandleGetStatsHasCorsHeaders(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(&request, mockWebsitStats{})
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
