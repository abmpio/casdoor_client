package sdk

import (
	"encoding/json"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

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
