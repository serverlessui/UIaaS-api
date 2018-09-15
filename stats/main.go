package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/stats/config"
	"github.com/serverlessui/UIaaS-api/stats/handler"
)

//GetStatistics is the entry point to the Lambda function
func GetStatistics(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.HandleGetStats(&request, config.CreateStatistics())
}

func main() {
	lambda.Start(GetStatistics)
}
