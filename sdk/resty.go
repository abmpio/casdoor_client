package sdk

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-resty/resty/v2"
)

type requestOptions struct {
	queryParams  url.Values
	bodyValue    interface{}
	formData     map[string]string
	accessToken  string
	ignoreResult bool
	// cookies
	cookies []*http.Cookie

	// basicAuth,
	clientId     string
	clientSecret string

	// headers
	headers map[string]string
}

type RequestOption func(*requestOptions)

// set basic Authorization
func RequestOptionWithBasicAuth(clientId string, clientSecret string) RequestOption {
	return func(ro *requestOptions) {
		ro.clientId = clientId
		ro.clientSecret = clientSecret
	}
}

// set Bearer Authorization
func RequestOptionWithAuthorization(accessToken string) RequestOption {
	return func(ro *requestOptions) {
		ro.accessToken = accessToken
	}
}

// set form data
func RequestOptionWithFormData(formData map[string]string) RequestOption {
	return func(ro *requestOptions) {
		ro.formData = formData
	}
}

// set cookie
func RequestOptionWithCookie(cookies []*http.Cookie) RequestOption {
	return func(ro *requestOptions) {
		ro.cookies = cookies
	}
}

func newRequestOptions() *requestOptions {
	r := &requestOptions{
		headers: map[string]string{
			"contentType":     "application/json",
			"Accept-Language": "zh",
		},
		queryParams: make(map[string][]string, 0),
		formData:    make(map[string]string),
	}
	return r
}
func doRequestWithResty(url, httpMethod string, opts ...RequestOption) (*resty.Response, error) {
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
	// set token(user accessToken)
	if len(options.accessToken) > 0 {
		r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", options.accessToken))
	}
	// set basic auth (client auth)
	if len(options.clientId) > 0 && len(options.clientSecret) > 0 {
		r.SetBasicAuth(options.clientId, options.clientSecret)
	}
	// queryParams
	if len(options.queryParams) > 0 {
		r.SetQueryParamsFromValues(options.queryParams)
	}
	// cookies
	if len(options.cookies) > 0 {
		r.SetCookies(options.cookies)
	}
	if len(options.formData) > 0 {
		r.SetFormData(options.formData)
	}
	// ignore response?
	if !options.ignoreResult {
		r.SetResult(&casdoorsdk.Response{})
		r.SetError(&casdoorsdk.Response{})
	}
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
