package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

//BucketInput is a struct which deines the required parameters to create an s3 bucket based site
type BucketInput struct {
	HostedZone        string `json:"hostedZone"`
	FullDomainName    string `json:"fullDomainName"`
	AcmCertificateArn string `json:"acmCertificateArn"`
	CacheValueTTL     string `json:"cacheValueTTL"`
	WebsiteURL        string `json:"websiteURL"`
}

//HandleCDN method to handle Bucket creation
func HandleCDN(request events.APIGatewayProxyRequest, action DeployAction) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	stackName := request.PathParameters[siteNameParam]

	var req BucketInput
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		fmt.Println("Received error: ", err)

		return events.APIGatewayProxyResponse{Body: "", StatusCode: 422}, nil
	}
	err = action.DeployBucket(&req, stackName)
	if err != nil {
		fmt.Println("Error deploying bucket: ", err)

		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 202}, nil
}
