package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/config"
	"github.com/serverlessui/UIaaS-api/handler"
)

//DeployHostedZone is the entry point to the Lambda function
func DeployHostedZone(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	action := config.CreateDeployAction()
	return handler.HandleHostedZone(request, action)
}

func main() {
	lambda.Start(DeployHostedZone)
}
