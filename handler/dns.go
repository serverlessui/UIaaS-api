package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

//DeployAction interface to define deploy action
type DeployAction interface {
	DeployBucket(bucketInput *BucketInput) error
	DeployHostedZone(dnsInput *DNSInput) (*Route53Output, error)
}

//Route53Output struct containing output from Route53
type Route53Output struct {
	WebsiteArn string
}

//DNSInput is a struct representing the required parameters to pass for HostedZoneCreation creation
type DNSInput struct {
	HostedZone string `json:"hostedZone"`
	//todo- add type safety
	HostedZoneExists string `json:"hostedZoneExists"`
	Environment      string `json:"environment"`
}

//HandleHostedZone handles Http requests
func HandleHostedZone(request events.APIGatewayProxyRequest, action DeployAction) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	var req DNSInput
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		fmt.Println("Received error: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 422}, nil
	}
	output, err := action.DeployHostedZone(&req)
	if err != nil {
		fmt.Println("Error deploying hosted zone: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	resp, err := json.Marshal(output)
	if err != nil {
		fmt.Println("Error marshalling value: ", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 202}, nil
}
