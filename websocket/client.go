package websocket

import (
	"encoding/json"
	"github.com/lxzan/gws"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) Connect() {
	opt := &gws.ClientOption{
		Addr: c.url,
	}
	client, _, err := gws.NewClient(&Handler{}, opt)
	if err != nil {
		logrus.Error(err)
		return
	}
	client.ReadLoop()
}

type Handler struct{}

func (h *Handler) OnOpen(socket *gws.Conn) {
	logrus.Infof("已连接至 WebSocket 网关 %s", socket.RemoteAddr())

	c := cron.New()
	// 每 25 秒发送一次心跳
	_, _ = c.AddFunc("@every 25s", func() {
		payload := &Payload{S: 2, SN: sn}
		message, _ := json.Marshal(payload)
		logrus.Debugf("发送心跳: %s", message)
		err := socket.WriteMessage(1, message)
		if err != nil {
			return
		}
	})
	c.Start()
}

func (h *Handler) OnClose(socket *gws.Conn, err error) {
	logrus.Infof("已从 WebSocket 网关 %s 断开连接, %+v", socket.RemoteAddr(), err)
}

func (h *Handler) OnPing(*gws.Conn, []byte) {}

func (h *Handler) OnPong(*gws.Conn, []byte) {}

func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	logrus.Debugf("接收消息: %s", message.Data)
	var payload Payload
	err := json.Unmarshal(message.Bytes(), &payload)
	if err != nil {
		logrus.Error(err)
		return
	}

	onPayload(&payload)
}
