package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/abmpio/configurationx"
	optCasdoor "github.com/abmpio/configurationx/options/casdoor"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-resty/resty/v2"
)

type ClientX struct {
	*casdoorsdk.Client
}

type httpClientWithLang struct {
	http.Client

	Lang string
}

var (
	_global_clientx      *ClientX
	clientx_instanceOnce sync.Once
)

// 获取ServiceGroup实例
func GetGlobalClientX() *ClientX {
	if _global_clientx != nil {
		return _global_clientx
	}
	clientx_instanceOnce.Do(func() {
		_global_clientx = NewCassdorClientXFromGlobal()
	})
	return _global_clientx
}

func (c *httpClientWithLang) Do(req *http.Request) (*http.Response, error) {
	if len(c.Lang) > 0 {
		req.Header.Set("Accept-Language", c.Lang)
	}
	return c.Client.Do(req)
}

func InitCasdoorsdkWithLang(lang string) {
	if len(lang) > 0 {
		casdoorsdk.SetHttpClient(&httpClientWithLang{
			Lang: lang,
		})
	}
}

func NewCassdorClientX(config *casdoorsdk.AuthConfig) *ClientX {
	x := &ClientX{}
	client := casdoorsdk.NewClientWithConf(config)
	x.Client = client
	return x
}

func GetSSOApplicationName() string {
	casdoorOpt := &optCasdoor.CasdoorOptions{}
	configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
	return casdoorOpt.ApplicationName
}

func GetOrganizationName() string {
	casdoorOpt := &optCasdoor.CasdoorOptions{}
	configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
	return casdoorOpt.OrganizationName
}

func NewCassdorClientXFromGlobal() *ClientX {
	casdoorOpt := &optCasdoor.CasdoorOptions{}
	configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
	authConfig := &casdoorsdk.AuthConfig{
		Endpoint:         casdoorOpt.Endpoint,
		ClientId:         casdoorOpt.ClientId,
		ClientSecret:     casdoorOpt.ClientSecret,
		Certificate:      casdoorOpt.Certificate,
		OrganizationName: casdoorOpt.OrganizationName,
		ApplicationName:  casdoorOpt.ApplicationName,
	}
	return NewCassdorClientX(authConfig)
}

// get organization list by owner
func (x *ClientX) GetOrganizationsByOwner(owner string) ([]*casdoorsdk.Organization, error) {
	queryMap := map[string]string{
		"owner": owner,
	}

	url := x.GetUrl("get-organizations", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var organizations []*casdoorsdk.Organization
	err = json.Unmarshal(bytes, &organizations)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

// get organization list by owner
func (x *ClientX) GetRolesByOwner(owner string) ([]*casdoorsdk.Role, error) {
	queryMap := map[string]string{
		"owner": owner,
	}

	url := x.GetUrl("get-roles", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var roles []*casdoorsdk.Role
	err = json.Unmarshal(bytes, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil

}

type requestOptions struct {
	queryParams url.Values
	bodyValue   interface{}
	formData    map[string]string
	accessToken string
}

func (c *ClientX) doRequestWithResty(url, httpMethod string, opts ...func(o *requestOptions)) (*resty.Response, error) {
	client := resty.New()
	r := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh")

	requestOptions := &requestOptions{}
	for _, eachOpt := range opts {
		eachOpt(requestOptions)
	}
	if len(requestOptions.accessToken) > 0 {
		r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", requestOptions.accessToken))
	}
	// queryParams
	if len(requestOptions.queryParams) > 0 {
		r.SetQueryParamsFromValues(requestOptions.queryParams)
	}
	if len(requestOptions.formData) > 0 {
		r.SetFormData(requestOptions.formData)
	}
	// ignore response?
	r.SetResult(&casdoorsdk.Response{})
	r.SetError(&casdoorsdk.Response{})
	if requestOptions.bodyValue != nil {
		r.SetBody(requestOptions.bodyValue)
	}
	resp, err := r.Execute(httpMethod, url)
	return resp, err
}

func (c *ClientX) unmarshalResponseValue(resp *resty.Response) (*casdoorsdk.Response, error) {
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
