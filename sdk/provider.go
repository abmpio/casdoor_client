package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

const (
	ProviderCategory_OAuth        = "OAuth"
	ProviderCategory_Email        = "Email"
	ProviderCategory_SMS          = "SMS"
	ProviderCategory_Storage      = "Storage"
	ProviderCategory_SAML         = "SAML"
	ProviderCategory_Payment      = "Payment"
	ProviderCategory_Captcha      = "Captcha"
	ProviderCategory_Notification = "Notification"
)

// get provider with secret
// owner: provider owner,can be admin, organization name...
// name: provider name
func (x *ClientX) GetProviderWithSecret(owner, name string) (*casdoorsdk.Provider, error) {
	queryMap := map[string]string{
		"id":         fmt.Sprintf("%s/%s", owner, name),
		"withSecret": "1",
		//使用admin用户
		"userId": fmt.Sprintf("%s/%s", x.OrganizationName, "admin"),
	}

	url := x.GetUrl("get-provider", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var provider *casdoorsdk.Provider
	err = json.Unmarshal(bytes, &provider)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

// get provider with secret
// provider: provider
func (x *ClientX) AddProviderWith(provider *casdoorsdk.Provider, optList ...func(*casdoorsdk.Provider)) (*casdoorsdk.Response, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", provider.Owner, provider.Name),
	}

	for _, eachOpt := range optList {
		eachOpt(provider)
	}
	postBytes, err := json.Marshal(provider)
	if err != nil {
		return nil, err
	}

	action := "add-provider"
	resp, err := x.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// update provider
// name: provider name
// provider: provider
func (x *ClientX) UpdateProviderWith(name string, provider *casdoorsdk.Provider, optList ...func(*casdoorsdk.Provider)) (*casdoorsdk.Response, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", provider.Owner, name),
	}

	for _, eachOpt := range optList {
		eachOpt(provider)
	}
	postBytes, err := json.Marshal(provider)
	if err != nil {
		return nil, err
	}

	action := "update-provider"
	resp, err := x.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return resp, err
	}
	if resp.Data != "Affected" {
		return resp, fmt.Errorf("无效的数据")
	}
	return resp, nil
}

// update provider
// name: provider name
// provider: provider
func (x *ClientX) DeleteProviderWith(provider *casdoorsdk.Provider, optList ...func(*casdoorsdk.Provider)) (*casdoorsdk.Response, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", provider.Owner, provider.Name),
	}

	for _, eachOpt := range optList {
		eachOpt(provider)
	}
	postBytes, err := json.Marshal(provider)
	if err != nil {
		return nil, err
	}

	action := "delete-provider"
	resp, err := x.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return resp, err
	}
	if resp.Data != "Affected" {
		return resp, fmt.Errorf("无效的数据")
	}
	return resp, nil
}
