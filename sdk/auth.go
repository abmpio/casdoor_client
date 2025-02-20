package sdk

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type WeixinQrcodeWebhookParameters struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
}

type WeixinQrcodeWebhookData struct {
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	FromUserName string `xml:"FromUserName"`
	Ticket       string `xml:"Ticket"`
}

func SendWeixinQrcodeWebhook(casdoorEndpoint string, p *WeixinQrcodeWebhookParameters, data *WeixinQrcodeWebhookData) (*resty.Response, error) {
	url := fmt.Sprintf("%s/api/webhook", casdoorEndpoint)
	response, err := doRequestWithResty(url, http.MethodPost, func(o *requestOptions) {
		if o.headers == nil {
			o.headers = map[string]string{}
		}
		o.headers[http.CanonicalHeaderKey("Content-Type")] = "application/xml"
		if o.queryParams == nil {
			o.queryParams = make(map[string][]string)
		}
		o.queryParams.Add("signature", p.Signature)
		o.queryParams.Add("timestamp", p.Timestamp)
		o.queryParams.Add("nonce", p.Nonce)

		o.bodyValue = data
	})
	return response, err
}
