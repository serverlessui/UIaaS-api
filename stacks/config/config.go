package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/larse514/aws-cloudformation-go"
	"github.com/serverlessui/UIaaS-api/stacks/actions"
	"github.com/serverlessui/UIaaS-api/stacks/bucket"
	"github.com/serverlessui/UIaaS-api/stacks/dns"
	"github.com/serverlessui/UIaaS-api/stacks/handler"
	"github.com/serverlessui/UIaaS-api/stacks/iaas"
)

//todo- refactor this to cache when retrieving dependency

//CreateDeployAction method to create implementation of DeployAction interface
func CreateDeployAction() handler.DeployAction {
	cloudformation := createCloudformation()
	return actions.ServerlessUI{DNS: CreateDNS(cloudformation), Bucket: CreateS3Bucket(cloudformation)}
}

//CreateDNS method to create implementation of DNS interface
func CreateDNS(cloudformation *cloudformation.CloudFormation) actions.DNS {
	return &dns.FullSite{Executor: createExecutor(cloudformation), Resource: createStack(cloudformation), IaaS: CreateIaaS()}
}

//CreateS3Bucket method to create implementation of Bucket interface
func CreateS3Bucket(cloudformation *cloudformation.CloudFormation) actions.Bucket {
	return &bucket.S3Bucket{Executor: createExecutor(cloudformation), Resource: createStack(cloudformation), IaaS: CreateIaaS()}
}

//CreateStack method to create implementation of Stack interface
func CreateStack() *cf.Stack {
	return createStack(createCloudformation())
}
func createExecutor(cloudformation *cloudformation.CloudFormation) cf.IaaSExecutor {
	return cf.IaaSExecutor{Client: cloudformation}
}

func createStack(cloudformation *cloudformation.CloudFormation) *cf.Stack {
	resource := cf.Stack{Client: cloudformation}
	return &resource
}

//CreateIaaS method to return implementation of Infrastructure interface
func CreateIaaS() iaas.Infrastructure {
	return &iaas.AWSTemplate{}
}

func createCloudformation() *cloudformation.CloudFormation {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(1)
	}

	return cloudformation.New(sess)
}
