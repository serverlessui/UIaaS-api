package statistics

import (
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

//AthenaStats is a struct
type AthenaStats struct {
	Client athenaiface.AthenaAPI
}

//Get method to retrieve statistics
func (stats AthenaStats) Get(sourceName string) (*athena.ResultSet, error) {
	//start query execution
	somequeryString := "SELECT count(distinct(requestip)) AS total FROM cloudfront_logs"
	sourceName = "default"
	input := athena.StartQueryExecutionInput{QueryExecutionContext: &athena.QueryExecutionContext{
		Database: &sourceName,
	},
		QueryString: &somequeryString,
	}

	output, err := stats.Client.StartQueryExecution(&input)

	if err != nil {
		return nil, err
	}

	resp, err := stats.Client.GetQueryResults(&athena.GetQueryResultsInput{
		QueryExecutionId: output.QueryExecutionId,
	})

	if err != nil {
		return nil, err
	}

	//profit
	return resp.ResultSet, nil
}
