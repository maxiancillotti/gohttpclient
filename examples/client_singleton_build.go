package examples

import (
	"httpClient"
	"time"
)

var (
	client = getHttpClient()
)

func getHttpClient() httpClient.HttpClient {
	client := httpClient.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeOut(3 * time.Second).
		Build()
	return client
}
