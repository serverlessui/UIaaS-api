package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverlessui/UIaaS-api/account/handler"
)

//CreateAccount is the entry point to the Lambda function
func GetAccount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.GetAccount(request)
}

func main() {
	lambda.Start(GetAccount)
}
