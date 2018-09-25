package handler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type mockWebsitStats struct {
}

func (mock mockWebsitStats) Get(sourceName string) (*VisitStatistics, error) {
	return &VisitStatistics{}, nil
}

type mockBadWebsitStats struct {
}

func (mock mockBadWebsitStats) Get(sourceName string) (*VisitStatistics, error) {
	return &VisitStatistics{}, errors.New("")
}
func TestHandleGetStats(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	res, err := HandleGetStats(&request, mockWebsitStats{})
	if err != nil {
		t.Log("Error encounterd when none expected ", err)
		t.Fail()
	}

	expected := "{\"total\":0,\"totalPerMonth\":[{\"data\":[0,5000,15000,8000,15000,9000,30000],\"label\":\"CCB Technologies\"}],\"totalLastWeek\":0,\"totalToday\":0,\"visitsByCountry\":null}"
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
