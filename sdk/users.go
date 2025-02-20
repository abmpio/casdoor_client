package sdk

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

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
