package websocket

import (
	"encoding/json"
	"github.com/lxzan/gws"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"zhenxin.me/kook/api"
)

var sn = 0

type Payload struct {
	S  int         `json:"s"`
	D  interface{} `json:"d"`
	SN int         `json:"sn"`
}

func onPayload(payload *Payload, token string, socket *gws.Conn) {
	switch payload.S {
	case 0:
		sn = payload.SN
		break
	case 1:
		onHello(payload.D, token, socket)
		break
	case 3:
		onPong()
		break
	case 5:
		break
	case 6:
		break
	}
}

type Hello struct {
	SessionId string `json:"session_id"`
}

func onHello(data interface{}, token string, socket *gws.Conn) {
	logrus.Debugf("收到 Hello 消息: %s", data)
	hello := &Hello{}
	err := mapstructure.Decode(data, hello)
	if err != nil {
		logrus.Error(err)
		return
	}

	botApi := api.NewApi(token)
	me, err := botApi.Me()
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("机器人 %s 已上线, 会话 ID: %s", me.Data.UserName, hello.SessionId)

	go func() {
		for {
			sleep := 30 + rand.Intn(10) - 5
			logrus.Debugf("等待 %d 秒后发送心跳", sleep)
			<-time.After(time.Duration(sleep) * time.Second)

			payload := &Payload{S: 2, SN: sn}
			message, _ := json.Marshal(payload)
			logrus.Debugf("发送心跳: %s", message)
			err := socket.WriteMessage(1, message)
			if err != nil {
				return
			}
		}
	}()
}

func onPong() {
	logrus.Debug("已收到心跳")
}
