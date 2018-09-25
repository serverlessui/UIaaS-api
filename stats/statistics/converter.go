package statistics

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

func convertSingleRowAthenaOutput(output *athena.GetQueryResultsOutput, input *handler.VisitStatistics) error {
	if len(output.ResultSet.Rows) == 0 {
		return nil
	}
	value, err := strconv.Atoi(*output.ResultSet.Rows[1].Data[0].VarCharValue)
	if err != nil {
		return err
	}
	input.Total = value

	return nil
}
