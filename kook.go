package kook

import (
	"github.com/sirupsen/logrus"
	"zhenxin.me/kook/api"
	"zhenxin.me/kook/internal"
	"zhenxin.me/kook/websocket"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Start() {
	internal.InitLogger()

	logrus.Info("正在获取网关地址...")
	botApi := api.NewApi(c.token)
	gateway, err := botApi.Gateway(false)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("正在连接到网关...")
	client := websocket.NewClient(gateway.Data.Url)
	client.Connect()
}
