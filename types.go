package geetest

type Ticket struct {
	LotNumber     string `json:"lot_number"`
	CaptchaOutput string `json:"captcha_output"`
	PassToken     string `json:"pass_token"`
	GenTime       uint   `json:"gen_time"`
}

type Success struct {
	Result      string         `json:"result"`
	Reason      string         `json:"reason"`
	CaptchaArgs map[string]any `json:"captcha_args"`
}

type Error struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}
