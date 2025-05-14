package sdk

import (
	"fmt"
	"strconv"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// get sms verifications list by owner
func (x *ClientX) GetSmsVerifications(owner string, pageIndex, pageSize int) (*casdoorsdk.Response, error) {
	queryMap := map[string]string{
		"owner":     owner,
		"p":         strconv.Itoa(pageIndex),
		"pageSize":  strconv.Itoa(pageSize),
		"field":     "",
		"value":     "",
		"sortField": "",
		"sortOrder": "",
	}

	url := x.GetUrl("get-verifications", queryMap)

	response, err := x.DoGetResponse(url)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf(response.Msg)
	}

	return response, nil
}
