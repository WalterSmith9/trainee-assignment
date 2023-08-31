package products

import (
	"fmt"
	"github.com/WalterSmith9/trainee-assignment/tests"
	"github.com/stretchr/testify/require"
	"testing"
)

var SegmentCreateRoute = "/segment"

var SegmentCreateResponseSuccess = `
{
	"code": 1,
	"message": "success"
}
`

func TestSegmentCreate(t *testing.T) {

	resp, err := tests.Client().R().
		EnableTrace().
		SetBody(map[string]string{
			"slug": "AVITO_VOICE_MESSAGES_TEST",
		}).
		SetHeader("Accept", "application/json").
		Post(fmt.Sprintf("%s%s", tests.URL, SegmentCreateRoute))
	if err != nil {
		t.Fail()
	}

	t.Log(resp.Status())
	if resp.Status() != "200 OK" {
		t.Log(fmt.Sprintf("bad status code! %s", resp.Status()))
		t.Fail()
	}

	require.JSONEq(t, SegmentCreateResponseSuccess, string(resp.Body()))
}
