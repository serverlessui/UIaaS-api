package statistics

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

type mockBadStartQueryAthenaClient struct {
	athenaiface.AthenaAPI
}

func (mock mockBadStartQueryAthenaClient) StartQueryExecution(*athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	return &athena.StartQueryExecutionOutput{}, errors.New("SOMEERROR")
}
func (mock mockBadStartQueryAthenaClient) GetQueryResults(*athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	return &athena.GetQueryResultsOutput{}, nil
}

type mockBadGetQueryAthenaClient struct {
	athenaiface.AthenaAPI
}

func (mock mockBadGetQueryAthenaClient) StartQueryExecution(*athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	return &athena.StartQueryExecutionOutput{}, nil
}
func (mock mockBadGetQueryAthenaClient) GetQueryResults(*athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	return &athena.GetQueryResultsOutput{}, errors.New("SOMEOTHERERROR")
}
func TestAthenaStatsWhenStartQueryFailsReturnsError(t *testing.T) {
	expected := "SOMEERROR"
	stats := AthenaStats{mockBadStartQueryAthenaClient{}}

	_, err := stats.Get("SOURCE")
	actual := err.Error()

	if actual != expected {
		t.Log("Expected error ", expected, " got ", actual)
		t.Fail()
	}
}

func TestAthenaStatsWhenGetQueryFailsReturnsError(t *testing.T) {
	expected := "SOMEOTHERERROR"
	stats := AthenaStats{mockBadGetQueryAthenaClient{}}

	_, err := stats.Get("SOURCE")
	actual := err.Error()

	if actual != expected {
		t.Log("Expected error ", expected, " got ", actual)
		t.Fail()
	}
}
