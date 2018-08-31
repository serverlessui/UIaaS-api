package handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type mockGoodAccountRepository struct {
}
type mockBadAccountRepository struct {
}
type mockEmptyAccount struct {
}

func (mock mockGoodAccountRepository) CreateAccount(account *Account) (*Account, error) {
	return &Account{AccountID: responseAccountID}, nil
}
func (mock mockGoodAccountRepository) GetAccount(accoundID string) (*Account, error) {
	return &Account{AccountID: responseAccountID}, nil

}
func (mock mockEmptyAccount) CreateAccount(account *Account) (*Account, error) {
	return &Account{AccountID: responseAccountID}, nil
}
func (mock mockEmptyAccount) GetAccount(accoundID string) (*Account, error) {
	return &Account{}, nil

}

func (mock mockBadAccountRepository) CreateAccount(account *Account) (*Account, error) {
	return &Account{}, errors.New("ERROR")

}
func (mock mockBadAccountRepository) GetAccount(accoundID string) (*Account, error) {
	return &Account{}, errors.New("ERROR")

}
func TestGetAccountStatusReturnsOK(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mock := mockGoodAccountRepository{}
	resp, err := GetAccount(request, mock)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := 200

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestGetAccountReturnsEmptyAccountIdReturnsNotFound(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mock := mockEmptyAccount{}
	resp, err := GetAccount(request, mock)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := 404

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
func TestGetAccountRepoReturnsErrorGetHandlerReturns503(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mock := mockBadAccountRepository{}
	resp, err := GetAccount(request, mock)

	if err != nil {
		t.Log("Encountered error when none was expected ", err)
		t.Fail()
	}

	actual := resp.StatusCode
	expected := http.StatusServiceUnavailable

	if expected != actual {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}
}
