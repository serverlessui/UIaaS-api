package statistics

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

func convertSingleRowAthenaOutput(output *athena.GetQueryResultsOutput, input *handler.VisitStatistics) error {

	for _, row := range output.ResultSet.Rows {
		if len(row.Data) == 0 {
			return nil
		}
		value, err := strconv.Atoi(*row.Data[1].VarCharValue)
		if err != nil {
			return err
		}
		input.Total = value
	}
	return nil
}
