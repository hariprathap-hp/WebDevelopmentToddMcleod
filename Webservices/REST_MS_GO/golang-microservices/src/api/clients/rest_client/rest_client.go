package rest_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	MockupsEnabled = false
	mocks          = make(map[string]*Mock)
)

type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

func StartMockups() {
	MockupsEnabled = true
}

func StopMockups() {
	MockupsEnabled = false
}

func AddMockups(mock Mock) {
	mocks[mock.URL] = &mock
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	if MockupsEnabled {
		mock := mocks[url]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
	}
	client := http.Client{}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header
	return client.Do(request)
}
