package sdk

import (
	"encoding/json"
	"fmt"
	"strconv"

	jsonutil "github.com/abmpio/libx/json"
)

type Captcha struct {
	Owner         string `json:"owner"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AppKey        string `json:"appKey"`
	Scene         string `json:"scene"`
	CaptchaId     string `json:"captchaId"`
	CaptchaImage  []byte `json:"captchaImage"`
	ClientId      string `json:"clientId"`
	ClientSecret  string `json:"clientSecret"`
	ClientId2     string `json:"clientId2"`
	ClientSecret2 string `json:"clientSecret2"`
	SubType       string `json:"subType"`
}

// get Captcha info
func (x *ClientX) GetCaptcha(owner string, applicationName string, isCurrentProvider bool) (*Captcha, error) {
	queryMap := map[string]string{
		"applicationId":     fmt.Sprintf("%s/%s", owner, applicationName),
		"isCurrentProvider": strconv.FormatBool(isCurrentProvider),
	}

	url := x.GetUrl("get-captcha", queryMap)

	bytes, err := x.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	captcha := Captcha{}
	err = json.Unmarshal(bytes, &captcha)
	if err != nil {
		return nil, err
	}
	return &captcha, nil
}

type VerificationForm struct {
	Dest          string `form:"dest" json:"dest"`
	Type          string `form:"type" json:"type"`
	CountryCode   string `form:"countryCode" json:"countryCode"`
	ApplicationId string `form:"applicationId" json:"applicationId"`
	Method        string `form:"method" json:"method"`
	CheckUser     string `form:"checkUser" json:"checkUser"`

	CaptchaType  string `form:"captchaType" json:"captchaType"`
	ClientSecret string `form:"clientSecret" json:"clientSecret"`
	CaptchaToken string `form:"captchaToken" json:"captchaToken"`
}

// send verification code
func (x *ClientX) SendVerificationCode(f *VerificationForm) (bool, error) {
	param := map[string]string{}
	err := jsonutil.ConvertObjectTo(f, &param)
	if err != nil {
		return false, err
	}

	bytes, err := json.Marshal(param)
	if err != nil {
		return false, err
	}

	resp, err := x.Client.DoPost("send-verification-code", nil, bytes, true, false)
	if err != nil {
		return false, err
	}

	return resp.Status == "ok", nil
}
