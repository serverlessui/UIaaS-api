package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/config"
	"github.com/serverlessui/UIaaS-api/handler"
)

//DeployBucketHandler is the entry point to the Lambda function
func DeployBucketHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	action := config.CreateDeployAction()
	return handler.HandleBucket(request, action)
}

func main() {
	lambda.Start(DeployBucketHandler)
}
