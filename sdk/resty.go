package sdk

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-resty/resty/v2"
)

type requestOptions struct {
	queryParams url.Values
	bodyValue   interface{}
	formData    map[string]string
	accessToken string

	// headers
	headers map[string]string
}

func newRequestOptions() *requestOptions {
	r := &requestOptions{
		headers: map[string]string{
			"contentType":     "application/json",
			"Accept-Language": "zh",
		},
	}
	return r
}
func doRequestWithResty(url, httpMethod string, opts ...func(o *requestOptions)) (*resty.Response, error) {
	client := resty.New()
	r := client.R().
		EnableTrace()

	options := newRequestOptions()
	for _, eachOpt := range opts {
		eachOpt(options)
	}

	// set headers
	if len(options.headers) > 0 {
		for eachKey, eachValue := range options.headers {
			if len(eachKey) > 0 {
				r.SetHeader(eachKey, eachValue)
			}
		}
	}
	// set token
	if len(options.accessToken) > 0 {
		r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", options.accessToken))
	}
	// queryParams
	if len(options.queryParams) > 0 {
		r.SetQueryParamsFromValues(options.queryParams)
	}
	if len(options.formData) > 0 {
		r.SetFormData(options.formData)
	}
	// ignore response?
	r.SetResult(&casdoorsdk.Response{})
	r.SetError(&casdoorsdk.Response{})
	if options.bodyValue != nil {
		r.SetBody(options.bodyValue)
	}
	resp, err := r.Execute(httpMethod, url)
	return resp, err
}

func unmarshalResponseValue(resp *resty.Response) (*casdoorsdk.Response, error) {
	if resp.IsSuccess() {
		casdoorRes, ok := resp.Result().(*casdoorsdk.Response)
		if ok {
			return casdoorRes, nil
		}
	} else {
		casdoorRes, ok := resp.Error().(*casdoorsdk.Response)
		if ok {
			return casdoorRes, nil
		}
	}
	body := string(resp.Body())
	return nil, errors.New(body)
}
