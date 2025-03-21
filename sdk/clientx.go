package sdk

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/abmpio/configurationx"
	optCasdoor "github.com/abmpio/configurationx/options/casdoor"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
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
	casdoorOpt := GetGlobalCasdoorOptions()
	return casdoorOpt.ApplicationName
}

func GetOrganizationName() string {
	casdoorOpt := GetGlobalCasdoorOptions()
	return casdoorOpt.OrganizationName
}

// get casdoor global options
func GetGlobalCasdoorOptions() *optCasdoor.CasdoorOptions {
	casdoorOpt := &optCasdoor.CasdoorOptions{}
	configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
	casdoorOpt.Normalize()
	return casdoorOpt
}

// create casdoorsdk.AuthConfig instance from optCasdoor.CasdoorOptions
func CasdoorAuthConfigFromCasdoorOptions(casdoorOpt *optCasdoor.CasdoorOptions) *casdoorsdk.AuthConfig {
	return &casdoorsdk.AuthConfig{
		Endpoint:         casdoorOpt.Endpoint,
		ClientId:         casdoorOpt.ClientId,
		ClientSecret:     casdoorOpt.ClientSecret,
		Certificate:      casdoorOpt.Certificate,
		OrganizationName: casdoorOpt.OrganizationName,
		ApplicationName:  casdoorOpt.ApplicationName,
	}
}

// 创建全局配置的casdoor client
func NewCassdorClientXFromGlobal() *ClientX {
	casdoorOpt := GetGlobalCasdoorOptions()
	return NewCassdorClientX(CasdoorAuthConfigFromCasdoorOptions(casdoorOpt))
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
