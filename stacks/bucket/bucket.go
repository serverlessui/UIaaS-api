package bucket

import (
	"log"

	"github.com/larse514/aws-cloudformation-go"
	"github.com/serverlessui/UIaaS-api/stacks/handler"
	"github.com/serverlessui/UIaaS-api/stacks/iaas"
)

const (
	//fullSite param values
	domainNameParam     = "HostedZone"
	fullDomainNameParam = "FullDomainName"
	acmCertARNParam     = "AcmCertificateArn"
	ttlCacheValueParam  = "CacheValueTTL"
	websiteURL          = "WebsiteBucketURL"
)

//S3Bucket is a struct to define needed S3 Bucket dependencies
type S3Bucket struct {
	Executor cf.Executor
	Resource cf.Resource
	IaaS     iaas.Infrastructure
}

//DeploySite is a function to Create an S3 Site with CDN and ACM
func (s3Bucket S3Bucket) DeploySite(input *handler.BucketInput, stackName string) error {
	stack, err := s3Bucket.Resource.GetStack(&stackName)
	if err != nil {
		return err
	}

	if *stack.StackName == "" {
		log.Println("Creating s3 bucket ", stackName)

		template, err := s3Bucket.IaaS.GetS3Site()

		if err != nil {
			return err
		}
		//create stack
		err = s3Bucket.Executor.CreateStack(*template, stackName, createInputParameters(input), nil)

		if err != nil {
			return err
		}
	}
	log.Println("S3 bucket already exists")
	return nil

}

//Helper method to create []*cloudformation.Parameter from input
func createInputParameters(input *handler.BucketInput) *map[string]string {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.HostedZone
	parameterMap[fullDomainNameParam] = input.FullDomainName
	parameterMap[acmCertARNParam] = input.AcmCertificateArn
	parameterMap[ttlCacheValueParam] = input.CacheValueTTL
	parameterMap[websiteURL] = input.WebsiteURL

	return &parameterMap

}
