package dns

import (
	"log"
	"strings"

	"github.com/larse514/aws-cloudformation-go"
	"github.com/serverlessui/UIaaS-api/handler"
	"github.com/serverlessui/UIaaS-api/iaas"
)

const (
	//FullSite param values
	domainNameParam       = "HostedZone"
	hostedZoneExistsParam = "HostedZoneExists"
	environmentParam      = "Environment"
	websiteArnOutput      = "WebsiteCertArn"
)

//FullSite is an implementation of the DNS interface
type FullSite struct {
	Executor cf.Executor
	Resource cf.Resource
	IaaS     iaas.Infrastructure
}

//DeployHostedZone Method to create fullSite hosted zone
func (fullSite FullSite) DeployHostedZone(input *handler.DNSInput) (*handler.Route53Output, error) {
	//replace domain name
	stackName := getStackName(input)

	//todo- i recommend refactoring this out of here
	stack, err := fullSite.Resource.GetStack(&stackName)
	if err != nil {
		return nil, err
	}
	websiteOutputValue := input.Environment + "-" + websiteArnOutput
	if *stack.StackName == "" {

		log.Println("Creating new dns stack")

		template, err := fullSite.IaaS.GetFullSite()

		if err != nil {
			return nil, err
		}
		//create stack
		err = fullSite.Executor.CreateStack(*template, stackName, createDNSInputParameters(input), nil)
		if err != nil {
			return nil, err
		}

		return &handler.Route53Output{}, nil

	}

	log.Println("DNS Stack already exists")
	return &handler.Route53Output{WebsiteArn: cf.GetOutputValue(stack, websiteOutputValue)}, nil
}

//Method to convert DomainName from input to stack name
//fullSite does not allow for full stop (.) characters
func getStackName(input *handler.DNSInput) string {
	return strings.Replace(input.HostedZone, ".", "-", -1)
}

//Helper method to create []*cloudformation.Parameter from input
func createDNSInputParameters(input *handler.DNSInput) *map[string]string {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.HostedZone
	parameterMap[environmentParam] = input.Environment
	parameterMap[hostedZoneExistsParam] = input.HostedZoneExists

	return &parameterMap

}

//NewRoute53 constructor for fullSite
func NewRoute53(executor cf.Executor, resource cf.Resource, iaas iaas.Infrastructure) *FullSite {
	return &FullSite{executor, resource, iaas}
}
