package sdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-resty/resty/v2"
)

type UserX struct {
	casdoorsdk.User

	MultiFactorAuths []*casdoorsdk.MfaProps `json:"multiFactorAuths,omitempty"`
}

// reset phone
func (x *ClientX) ResetPhone(accessToken string, code string, newPhone string) error {
	param := make(map[string]string, 0)
	param["type"] = "phone"
	param["dest"] = newPhone
	param["code"] = code

	url := x.GetUrl("reset-email-or-phone", nil)
	resp, err := doRequestWithResty(url, resty.MethodPost, func(o *requestOptions) {
		o.accessToken = accessToken
		o.formData = param
	})
	if err != nil {
		return err
	}
	result, err := unmarshalResponseValue(resp)
	if err != nil {
		return err
	}

	if result.Status != "ok" {
		return errors.New(result.Msg)
	}
	return nil
}

func (x *ClientX) ResetEmail(accessToken string, code string, newEmail string) error {
	param := make(map[string]string, 0)
	param["type"] = "email"
	param["dest"] = newEmail
	param["code"] = code

	url := x.GetUrl("reset-email-or-phone", nil)
	resp, err := doRequestWithResty(url, resty.MethodPost, func(o *requestOptions) {
		o.accessToken = accessToken
		o.formData = param
	})
	if err != nil {
		return err
	}
	result, err := unmarshalResponseValue(resp)
	if err != nil {
		return err
	}

	if result.Status != "ok" {
		return errors.New(result.Msg)
	}
	return nil
}

type IdpUserInfo struct {
	Id          string            `json:"id"`
	Username    string            `json:"username"`
	DisplayName string            `json:"displayName"`
	UnionId     string            `json:"unionId"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	CountryCode string            `json:"countryCode"`
	AvatarUrl   string            `json:"avatarUrl"`
	Extra       map[string]string `json:"extra"`
}

// link user to oauth
func (x *ClientX) LinkUserOAuth(organization, userName string,
	providerType string, idpUserInfo IdpUserInfo) error {

	postBytes, err := json.Marshal(map[string]interface{}{
		"providerType": providerType,
		"idpUserInfo":  idpUserInfo,
	})
	if err != nil {
		return err
	}

	resp, err := x.DoPost("link-user-oauth", map[string]string{
		"id": fmt.Sprintf("%s/%s", organization, userName),
	}, postBytes, false, false)
	if err != nil {
		return err
	}
	if resp.Status != "ok" {
		return errors.New(resp.Msg)
	}
	return nil
}

// unlink user to oauth
func (x *ClientX) UnlinkUserOAuth(user casdoorsdk.User, providerType string) error {
	postBytes, err := json.Marshal(map[string]interface{}{
		"providerType": providerType,
		"user":         user,
	})
	if err != nil {
		return err
	}

	resp, err := x.DoPost("unlink", map[string]string{
		"userId": fmt.Sprintf("%s/%s", user.Owner, user.Name),
	}, postBytes, false, false)
	if err != nil {
		return err
	}
	if resp.Status != "ok" {
		return errors.New(resp.Msg)
	}
	return nil
}

// get user by organization
// not paging
func (x *ClientX) GetUsersByOrganization(organization string) ([]*casdoorsdk.User, error) {
	queryMap := map[string]string{
		"owner": organization,
	}

	url := x.GetUrl("get-users", queryMap)
	response, err := x.DoGetResponse(url)
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var users []*casdoorsdk.User
	err = json.Unmarshal(dataBytes, &users)
	if err != nil {
		return nil, errors.New("response data format is incorrect")
	}

	return users, nil
}

// get user by field value
func (x *ClientX) GetUserByField(organization, field, value string) (*casdoorsdk.User, error) {
	queryMap := map[string]string{
		"field": field,
		"value": value,
	}

	// setup organization
	if len(organization) <= 0 {
		queryMap["organization"] = x.OrganizationName
	} else {
		queryMap["organization"] = organization
	}

	url := x.GetUrl("get-user-by-field", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *casdoorsdk.User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// get userx by organization,userId
func (x *ClientX) GetUserXByUserId(organization, userId string) (*UserX, error) {
	queryMap := map[string]string{
		"owner":  organization,
		"userId": userId,
	}

	url := x.GetUrl("get-user", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *UserX
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// get userx by organization,username
func (x *ClientX) GetEmailAndPhone(organization, userName string) (*UserX, error) {
	queryMap := map[string]string{
		"organization": organization,
		"username":     userName,
	}

	url := x.GetUrl("get-email-and-phone", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *UserX
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
