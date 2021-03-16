package bluefox

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Simply fetches from the botblock api in an easy way!
func Fetch(method string, url string, structure interface{}, token string, body map[string]string) error {
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}

	marshalledBody, _ := json.Marshal(body)
	jsonBody := bytes.NewBuffer(marshalledBody)

	req, err := http.NewRequest(method, BaseURL+url, jsonBody)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	res, err := client.Do(req)

	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return errors.New("UnexpectedError: Failed reading request data")
	}

	marshallErr := json.Unmarshal(data, structure)

	if marshallErr != nil {
		return errors.New("UnexpectedError: Failed while marshalling json: " + marshallErr.Error())
	}

	if res.StatusCode != 200 {
		return errors.New("BluefoxHostAPIError: Bluefox api sent an unusual api response as " + string(data) + " with status code as " + strconv.Itoa(res.StatusCode) + "!")
	}

	return nil
}
