package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/config"
	"github.com/serverlessui/UIaaS-api/handler"
)

//DeployCDNHandler is the entry point to the Lambda function
func DeployCDNHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	action := config.CreateDeployAction()
	return handler.HandleCDN(request, action)
}

func main() {
	lambda.Start(DeployCDNHandler)
}
