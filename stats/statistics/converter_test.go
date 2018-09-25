package statistics

import (
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

func TestWhenTotalIsReceivedGetQueryResultsOutputContainsTotal(t *testing.T) {
	totalVarCharValue := "total"
	totalValue := "2"
	input := handler.VisitStatistics{}
	output := &athena.GetQueryResultsOutput{
		ResultSet: &athena.ResultSet{
			Rows: []*athena.Row{&athena.Row{
				Data: []*athena.Datum{
					&athena.Datum{VarCharValue: &totalVarCharValue},
					&athena.Datum{VarCharValue: &totalValue},
				},
			}},
		},
	}

	err := convertSingleRowAthenaOutput(output, &input)
	if err != nil {
		t.Log("Error received when none expected ", err)
		t.Fail()
	}
	expected, _ := strconv.Atoi(totalValue)
	actual := input.Total

	if actual != expected {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestWhenTotalIsInvalidGetQueryResultsOutputReturnsError(t *testing.T) {
	totalVarCharValue := "total"
	totalValue := "NOT AN INT"
	input := handler.VisitStatistics{}
	output := &athena.GetQueryResultsOutput{
		ResultSet: &athena.ResultSet{
			Rows: []*athena.Row{&athena.Row{
				Data: []*athena.Datum{
					&athena.Datum{VarCharValue: &totalVarCharValue},
					&athena.Datum{VarCharValue: &totalValue},
				},
			}},
		},
	}

	err := convertSingleRowAthenaOutput(output, &input)
	if err == nil {
		t.Log("Error not received when one expected ")
		t.Fail()
	}

}

func TestWhenNoResultsReturnedGetQueryResultsOutputReturns0Total(t *testing.T) {
	input := handler.VisitStatistics{}
	output := &athena.GetQueryResultsOutput{
		ResultSet: &athena.ResultSet{
			Rows: []*athena.Row{&athena.Row{
				Data: []*athena.Datum{},
			}},
		},
	}

	err := convertSingleRowAthenaOutput(output, &input)
	if err != nil {
		t.Log("Error received when none expected ", err)
		t.Fail()
	}
	expected := 0
	actual := input.Total

	if actual != expected {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}
