package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/stacks/config"
	"github.com/serverlessui/UIaaS-api/stacks/handler"
)

//GetStack is the entry point to the Lambda function
func GetStack(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	action := config.CreateStack()
	return handler.HandleGetStack(request, action)
}

func main() {
	lambda.Start(GetStack)
}
