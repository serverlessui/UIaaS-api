package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/stats/config"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

//GetStatisticsMeta is the entry point to the Lambda function
func GetStatisticsMeta(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.HandleGetStatsMeta(&request, config.CreateRepository())
}

func main() {
	lambda.Start(GetStatisticsMeta)
}
