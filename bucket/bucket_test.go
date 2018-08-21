package bucket

import (
	"testing"

	"github.com/serverlessui/UIaaS-api/handler"
)

const (
	domainName       = "somedomain.com"
	longerDomainName = "prefix.somedomain.com"
)

func TestGetStackName(t *testing.T) {
	input := handler.BucketInput{FullDomainName: domainName}
	expected := "somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}

func TestGetStackNamePrefix(t *testing.T) {
	input := handler.BucketInput{FullDomainName: longerDomainName}

	expected := "prefix-somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}
