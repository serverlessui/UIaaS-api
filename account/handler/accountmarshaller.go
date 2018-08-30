package handler

import "encoding/json"

//MarshallAccountStructToString method to marshall account to a string
func MarshallAccountStructToString(account *Account) (string, error) {
	resp, err := json.Marshal(account)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
