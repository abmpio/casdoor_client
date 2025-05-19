package sdk

import (
	"encoding/json"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type MfaSetupInitiateForm struct {
	Owner string `form:"owner" json:"owner"`
	// user name
	Name    string `form:"name" json:"name"`
	MfaType string `form:"mfaType" json:"mfaType"`
}

// setup MFA
func (x *ClientX) MfaSetupInitiate(f *MfaSetupInitiateForm) (*casdoorsdk.Response, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}

	resp, err := x.Client.DoPost("/mfa/setup/initiate", nil, bytes, true, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MfaSetupVerifyForm struct {
	Owner       string `form:"owner" json:"owner"`
	Name        string `form:"name" json:"name"`
	MfaType     string `form:"mfaType" json:"mfaType"`
	Passcode    string `form:"passcode" json:"passcode"`
	Secret      string `form:"secret" json:"secret"`
	Dest        string `form:"dest" json:"dest"`
	CountryCode string `form:"countryCode" json:"countryCode"`
}

// setup MFA verify
func (x *ClientX) MfaSetupVerify(f *MfaSetupVerifyForm) (*casdoorsdk.Response, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}

	resp, err := x.Client.DoPost("/mfa/setup/verify", nil, bytes, true, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MfaSetupEnableForm struct {
	Owner         string `form:"owner" json:"owner"`
	Name          string `form:"name" json:"name"`
	MfaType       string `form:"mfaType" json:"mfaType"`
	Secret        string `form:"secret" json:"secret"`
	Dest          string `form:"dest" json:"dest"`
	CountryCode   string `form:"countryCode" json:"countryCode"`
	RecoveryCodes string `form:"recoveryCodes" json:"recoveryCodes"`
}

// setup MFA enable
func (x *ClientX) MfaSetupEnable(f *MfaSetupEnableForm) (*casdoorsdk.Response, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}

	resp, err := x.Client.DoPost("/mfa/setup/enable", nil, bytes, true, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MfaDeleteForm struct {
	Owner string `form:"owner" json:"owner"`
	// user name
	Name string `form:"name" json:"name"`
}

// delete MFA
func (x *ClientX) MfaDelete(input *MfaDeleteForm) (*casdoorsdk.Response, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := x.Client.DoPost("/delete-mfa", nil, bytes, true, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MfaSetPreferredForm struct {
	Owner string `form:"owner" json:"owner"`
	// user name
	Name    string `form:"name" json:"name"`
	MfaType string `form:"mfaType" json:"mfaType"`
}

// set preferred mfa
func (x *ClientX) MfaSetPreferred(input *MfaSetPreferredForm) (*casdoorsdk.Response, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := x.Client.DoPost("/set-preferred-mfa", nil, bytes, true, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
