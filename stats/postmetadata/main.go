package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/stats/config"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

//PostStatisticsMeta is the entry point to the Lambda function
func PostStatisticsMeta(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.HandlePostStatsMeta(&request, config.CreateRepository())
}

func main() {
	lambda.Start(PostStatisticsMeta)
}
