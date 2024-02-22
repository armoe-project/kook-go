package api

import (
	"encoding/json"
	"zhenxin.me/kook/internal/request"
)

type Gateway struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Url string `json:"url"`
	} `json:"data"`
}

func (a *Api) Gateway(compress bool) (*Gateway, error) {
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

	var gateway Gateway
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
