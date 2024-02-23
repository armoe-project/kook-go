package websocket

import (
	"encoding/json"
	"github.com/lxzan/gws"
	"github.com/sirupsen/logrus"
)

type Client struct {
	token string
	url   string
}

func NewClient(url, token string) *Client {
	return &Client{url: url, token: token}
}

func (c *Client) Connect() {
	opt := &gws.ClientOption{
		Addr: c.url,
	}
	client, _, err := gws.NewClient(&Handler{token: c.token}, opt)
	if err != nil {
		logrus.Error(err)
		return
	}
	client.ReadLoop()
}

type Handler struct {
	token string
}

func (h *Handler) OnOpen(socket *gws.Conn) {
	logrus.Infof("已连接至 WebSocket 网关 %s", socket.RemoteAddr())
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

	onPayload(&payload, h.token, socket)
}
