package sdk

import (
	"fmt"
	"net/http"

	jsonUtil "github.com/abmpio/libx/json"
	"github.com/abmpio/libx/slicex"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-resty/resty/v2"
)

// get application signup item by key
func GetSignupItem(app *casdoorsdk.Application, key string) *casdoorsdk.SignupItem {
	if len(app.SignupItems) <= 0 {
		return nil
	}
	return slicex.FindOne(app.SignupItems, func(item *casdoorsdk.SignupItem) bool {
		return item.Name == key
	})
}

// get application with cookie
func (x *ClientX) GetApplicationWithCookie(name string, opts ...RequestOption) (*http.Response, *casdoorsdk.Application, error) {
	url := x.GetUrl("get-application", nil)

	if len(opts) <= 0 {
		opts = make([]RequestOption, 0)
	}
	opts = append(opts, func(o *requestOptions) {
		o.queryParams.Add("id", fmt.Sprintf("%s/%s", "admin", name))
	})
	opts = append(opts, RequestOptionWithBasicAuth(x.ClientId, x.ClientSecret))

	resp, err := doRequestWithResty(url, resty.MethodGet, opts...)
	if err != nil {
		return resp.RawResponse, nil, err
	}

	result, err := unmarshalResponseValue(resp)
	if err != nil {
		return resp.RawResponse, nil, err
	}
	if result.Data == nil {
		return resp.RawResponse, nil, nil
	}

	application := &casdoorsdk.Application{}
	err = jsonUtil.ConvertObjectTo(result.Data, application)
	if err != nil {
		return resp.RawResponse, nil, err
	}
	return resp.RawResponse, application, nil
}
