package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpClientMock struct{}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {

	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	if mock := MockupServer.getMock(request.Method, request.URL.String(), string(body)); mock != nil {
		return mock.GetResponse(request)
	}
	return nil, fmt.Errorf("error retrieving mock")
}
