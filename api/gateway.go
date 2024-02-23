package api

import (
	"encoding/json"
	"zhenxin.me/kook/internal/request"
)

type GatewayIndex struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Url string `json:"url"`
	} `json:"data"`
}

func (a *Api) GatewayIndex(compress bool) (*GatewayIndex, error) {
	path := "/gateway/index"
	if compress {
		path += "?compress=1"
	} else {
		path += "?compress=0"
	}

	client := request.NewClient(a.token)
	body, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	var gateway GatewayIndex
	err = json.Unmarshal([]byte(body), &gateway)
	if err != nil {
		return nil, err
	}

	if gateway.Code != 0 {
		err = &Error{Code: gateway.Code, Message: gateway.Message}
		return nil, err
	}

	return &gateway, nil
}

type GatewayVoice struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		GatewayUrl string `json:"gateway_url"`
	}
}

func (a *Api) GatewayVoice(channelId string) (*GatewayVoice, error) {
	path := "/gateway/voice?channel_id=" + channelId

	client := request.NewClient(a.token)
	body, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	var gateway GatewayVoice
	err = json.Unmarshal([]byte(body), &gateway)
	if err != nil {
		return nil, err
	}

	if gateway.Code != 0 {
		err = &Error{Code: gateway.Code, Message: gateway.Message}
		return nil, err
	}

	return &gateway, nil
}
