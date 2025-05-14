package sdk

import (
	"encoding/json"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

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
