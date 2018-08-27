package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/larse514/aws-cloudformation-go"
)

const (
	stackNameParam = "stackname"
)

//HandleGetStack method to handle retrieval of stack resoruce
func HandleGetStack(request events.APIGatewayProxyRequest, cf *cf.Stack) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received parameters: ", request.PathParameters)

	stackName := request.PathParameters[stackNameParam]

	if stackName == "" {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 422}, nil

	}
	stack, err := cf.GetStack(&stackName)

	if err != nil {
		fmt.Println("error processing request: ", err)
	}

	if *stack.StackName == "" {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 404}, nil
	}

	resp, err := json.Marshal(stack)

	if err != nil {
		fmt.Println("Error marshalling value: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
}
