package actions

import (
	"log"

	"github.com/serverlessui/UIaaS-api/stacks/handler"
)

//Bucket is an interface to define creation of Bucket based sites
type Bucket interface {
	DeploySite(input *handler.BucketInput, stackName string) error
}

//DNS is an interface to represent Cloud DNS Services
type DNS interface {
	DeployHostedZone(input *handler.DNSInput, stackName string) (*handler.Route53Output, error)
}

//ServerlessUI struct to implement DeployAction
type ServerlessUI struct {
	DNS    DNS
	Bucket Bucket
}

//DeployBucket method to deploy serverless UI
func (serverless ServerlessUI) DeployBucket(bucketInput *handler.BucketInput, stackName string) error {
	err := serverless.Bucket.DeploySite(bucketInput, stackName)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		return err
	}
	return nil
}

//DeployHostedZone method to deploy serverless hosted zone
func (serverless ServerlessUI) DeployHostedZone(dnsInput *handler.DNSInput, stackName string) (*handler.Route53Output, error) {
	output, err := serverless.DNS.DeployHostedZone(dnsInput, stackName)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		return &handler.Route53Output{}, err
	}
	return output, nil
}
