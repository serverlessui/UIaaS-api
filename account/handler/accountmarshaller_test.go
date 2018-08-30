package handler

import "testing"

func TestMarshallAccountStructToString(t *testing.T) {
	expected := "{ACCOUNTID NAME     }"
	account := Account{
		AccountID: "ACCOUNTID",
		Name:      "NAME",
	}

	actual, err := MarshallAccountStructToString(&account)

	if err != nil {
		t.Log("Got error when none expected ", err)
		t.Fail()
	}

	if actual == "" {
		t.Log("Expected ", expected, " got ", account)
		t.Fail()
	}
}
