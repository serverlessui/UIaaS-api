package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/account/handler"
)

//CreateAccount is the entry point to the Lambda function
func CreateAccount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.CreateAccount(request)
}

func main() {
	lambda.Start(CreateAccount)
}
