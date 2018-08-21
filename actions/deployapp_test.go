package actions

import (
	"errors"
	"testing"

	"github.com/serverlessui/UIaaS-api/handler"
)

type mockBucket struct {
}

func (mock mockBucket) DeploySite(input *handler.BucketInput) error {
	return nil
}

type mockBadBucket struct {
}

func (mock mockBadBucket) DeploySite(input *handler.BucketInput) error {
	return errors.New("")
}

type mockDNS struct {
}

func (mock mockDNS) DeployHostedZone(input *handler.DNSInput) (*handler.Route53Output, error) {
	return &handler.Route53Output{WebsiteArn: "SOMEARN"}, nil
}

type mockBadDNS struct {
}

func (mock mockBadDNS) DeployHostedZone(input *handler.DNSInput) (*handler.Route53Output, error) {
	return nil, errors.New("")
}

func TestServerlessUIDeployBucketOk(t *testing.T) {
	ui := ServerlessUI{mockDNS{}, mockBucket{}}

	err := ui.DeployBucket(&handler.BucketInput{})

	if err != nil {
		t.Log("error encountered when none expected ", err)
		t.Fail()
	}
}
func TestServerlessUIDeployBucketReturnsError(t *testing.T) {
	ui := ServerlessUI{mockDNS{}, mockBadBucket{}}

	err := ui.DeployBucket(&handler.BucketInput{})

	if err == nil {
		t.Log("error encountered one expected ")
		t.Fail()
	}
}

func TestServerlessUIDeployHostedZoneOk(t *testing.T) {
	ui := ServerlessUI{mockDNS{}, mockBucket{}}

	_, err := ui.DeployHostedZone(&handler.DNSInput{})

	if err != nil {
		t.Log("error encountered when none expected ", err)
		t.Fail()
	}
}
func TestServerlessUIDeployHostedZoneReturnsError(t *testing.T) {
	ui := ServerlessUI{mockBadDNS{}, mockBucket{}}

	_, err := ui.DeployHostedZone(&handler.DNSInput{})

	if err == nil {
		t.Log("error encountered one expected ")
		t.Fail()
	}
}
