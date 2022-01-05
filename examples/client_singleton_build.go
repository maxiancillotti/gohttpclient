package examples

import (
	"time"

	"github.com/maxiancillotti/gohttpclient"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttpclient.Client {
	client := gohttpclient.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
