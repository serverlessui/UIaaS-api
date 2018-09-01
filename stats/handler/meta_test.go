package handler

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

const (
	accountID = "SOMEID"
	sid       = "SITEID"
)

type mockMetaDataRepository struct {
}

func (mock mockMetaDataRepository) CreateMetaData(metadatum *StatsMetaData) error {
	return nil
}

func (mock mockMetaDataRepository) GetMetaData(siteID string) (*StatsMetaData, error) {
	if siteID != sid {
		return &StatsMetaData{}, errors.New("ERROR")
	}
	return &StatsMetaData{AccountID: accountID}, nil
}

type mockBadMetaDataRepository struct {
}

func (mock mockBadMetaDataRepository) CreateMetaData(metadatum *StatsMetaData) error {
	return errors.New("SOME ERROR")
}

func (mock mockBadMetaDataRepository) GetMetaData(siteID string) (*StatsMetaData, error) {
	return &StatsMetaData{}, errors.New("ERROR")
}

func TestHandlePostStatsMetaReturns201WhenRepositoryReturnsNilError(t *testing.T) {
	request := events.APIGatewayProxyRequest{Body: "{}"}
	repo := mockMetaDataRepository{}
	response, err := HandlePostStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := 201
	actual := response.StatusCode

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestHandlePostStatsMetaReturns503WhenRepositoryReturnsError(t *testing.T) {
	request := events.APIGatewayProxyRequest{Body: "{}"}
	repo := mockBadMetaDataRepository{}
	response, err := HandlePostStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := 503
	actual := response.StatusCode

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestHandleGetStatsMetaReturns50PUnmarshallErrorOccurs(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	repo := mockBadMetaDataRepository{}
	response, err := HandlePostStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := 500
	actual := response.StatusCode

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestHandleGetStatsMetaReturns200WhenRepositoryReturnsNilError(t *testing.T) {
	pathParams := make(map[string]string)

	pathParams["sid"] = sid
	request := events.APIGatewayProxyRequest{PathParameters: pathParams}

	repo := mockMetaDataRepository{}
	response, err := HandleGetStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := 200
	actual := response.StatusCode

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestHandleGetStatsMetaReturnsRepositorysResponse(t *testing.T) {
	pathParams := make(map[string]string)

	pathParams["sid"] = sid
	request := events.APIGatewayProxyRequest{PathParameters: pathParams}

	repo := mockMetaDataRepository{}
	response, err := HandleGetStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := "{\"accountID\":\"SOMEID\",\"siteID\":\"\",\"tableName\":\"\",\"createdTimestamp\":0}"
	actual := response.Body

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}

func TestHandleProvidesSiteIdToMetaDataRepository(t *testing.T) {
	pathParams := make(map[string]string)

	pathParams["sid"] = sid
	request := events.APIGatewayProxyRequest{PathParameters: pathParams}
	repo := mockMetaDataRepository{}
	response, err := HandleGetStatsMeta(&request, repo)

	if err != nil {
		t.Log("Error encountered when none expected ", err)
		t.Fail()
	}
	expected := 200
	actual := response.StatusCode

	if expected != actual {
		t.Log("Got ", actual, " expected ", expected)
		t.Fail()
	}

}
