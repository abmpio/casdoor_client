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
	_global_clientx              *ClientX
	_global_clientx_instanceOnce sync.Once

	// built-in admin client
	_builtInAdminClientx               *ClientX
	_builtInAdminClientx_instanceOnece sync.Once
)

// 获取ServiceGroup实例
func GetGlobalClientX() *ClientX {
	if _global_clientx != nil {
		return _global_clientx
	}
	_builtInAdminClientx_instanceOnece.Do(func() {
		_global_clientx = NewCassdorClientXFromGlobal()
	})
	return _global_clientx
}

// get built-in admin clientx
func GetBuiltInAdminClientX() *ClientX {
	if _builtInAdminClientx != nil {
		return _builtInAdminClientx
	}
	_global_clientx_instanceOnce.Do(func() {
		_builtInAdminClientx = NewCassdorClientX(CasdoorAuthConfigFromCasdoorOptions(GetGlobalCasdoorOptions(), true))
	})
	return _builtInAdminClientx
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

// get application name
func GetSSOApplicationName() string {
	casdoorOpt := GetGlobalCasdoorOptions()
	return casdoorOpt.ApplicationName
}

// get organization name
func GetOrganizationName() string {
	casdoorOpt := GetGlobalCasdoorOptions()
	return casdoorOpt.OrganizationName
}

// get casdoor global options
func GetGlobalCasdoorOptions() *optCasdoor.CasdoorOptions {
	return configurationx.GetInstance().Casdoor
	// casdoorOpt := &optCasdoor.CasdoorOptions{}
	// configurationx.GetInstance().UnmarshalPropertiesTo(optCasdoor.ConfigurationKey, casdoorOpt)
	// casdoorOpt.Normalize()
	// return casdoorOpt
}

// create casdoorsdk.AuthConfig instance from optCasdoor.CasdoorOptions
func CasdoorAuthConfigFromCasdoorOptions(casdoorOpt *optCasdoor.CasdoorOptions, isBuiltIn bool) *casdoorsdk.AuthConfig {
	if !isBuiltIn {
		return &casdoorsdk.AuthConfig{
			Endpoint:         casdoorOpt.Endpoint,
			ClientId:         casdoorOpt.ClientId,
			ClientSecret:     casdoorOpt.ClientSecret,
			Certificate:      casdoorOpt.Certificate,
			OrganizationName: casdoorOpt.OrganizationName,
			ApplicationName:  casdoorOpt.ApplicationName,
		}
	} else {
		return &casdoorsdk.AuthConfig{
			Endpoint:         casdoorOpt.Endpoint,
			ClientId:         casdoorOpt.AppBuiltinClientId,
			ClientSecret:     casdoorOpt.AppBuiltinClientSecret,
			Certificate:      casdoorOpt.AppBuiltinCertificate,
			OrganizationName: BuiltInOrganization,
			ApplicationName:  AppBuiltIn,
		}
	}
}

// 创建全局配置的casdoor client
func NewCassdorClientXFromGlobal() *ClientX {
	casdoorOpt := GetGlobalCasdoorOptions()
	return NewCassdorClientX(CasdoorAuthConfigFromCasdoorOptions(casdoorOpt, false))
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
