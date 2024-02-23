package api

import (
	"encoding/json"
	"zhenxin.me/kook/internal/request"
)

type Me struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id             string `json:"id"`
		UserName       string `json:"username"`
		IdentifyNum    string `json:"identify_num"`
		Online         bool   `json:"online"`
		OS             string `json:"os"`
		Status         int    `json:"status"`
		Avatar         string `json:"avatar"`
		Banner         string `json:"banner"`
		Bot            bool   `json:"bot"`
		MobileVerified bool   `json:"mobile_verified"`
		MobilePrefix   string `json:"mobile_prefix"`
		Mobile         string `json:"mobile"`
		InvitedCount   int    `json:"invited_count"`
	}
}

func (a *Api) Me() (*Me, error) {
	path := "/user/me"

	client := request.NewClient(a.token)
	body, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	var me Me
	err = json.Unmarshal([]byte(body), &me)
	if err != nil {
		return nil, err
	}

	if me.Code != 0 {
		err = &Error{Code: me.Code, Message: me.Message}
		return nil, err
	}

	return &me, nil
}
