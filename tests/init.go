package tests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

var (
	URL = "http://api:8070/api/v2"
)

func Client() *resty.Client {
	client := resty.New()
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	return client
}
