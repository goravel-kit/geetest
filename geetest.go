package geetest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/support/json"
)

type Geetest struct {
	client     *resty.Client
	CaptchaID  string
	CaptchaKey string
}

func NewGeetest(config config.Config) *Geetest {
	client := resty.New()
	client.SetBaseURL(config.GetString("geetest.api_url"))
	client.SetTimeout(5 * time.Second)
	client.SetJSONMarshaler(json.Marshal)
	client.SetJSONUnmarshaler(json.Unmarshal)

	return &Geetest{
		client:     client,
		CaptchaID:  config.GetString("geetest.captcha_id"),
		CaptchaKey: config.GetString("geetest.captcha_key"),
	}
}

func (r *Geetest) Verify(ticket Ticket) (bool, error) {
	resp, err := r.client.R().
		SetFormData(map[string]string{
			"lot_number":     ticket.LotNumber,
			"captcha_output": ticket.CaptchaOutput,
			"pass_token":     ticket.PassToken,
			"gen_time":       fmt.Sprintf("%d", ticket.GenTime),
			"sign_token":     hmacEncode(r.CaptchaKey, ticket.LotNumber),
		}).
		SetQueryParam("captcha_id", r.CaptchaID).
		SetResult(&Success{}).
		SetError(&Error{}).
		Post("/validate")

	if err != nil {
		return false, err
	}
	if !resp.IsSuccess() {
		return false, fmt.Errorf("%s %s", resp.Status(), resp.Error().(*Error).Msg)
	}

	res := resp.Result().(*Success)
	if res.Result == "success" {
		return true, nil
	} else {
		return false, fmt.Errorf("%s", res.Reason)
	}
}

func hmacEncode(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
