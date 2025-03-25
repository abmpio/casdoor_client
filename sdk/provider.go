package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
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
