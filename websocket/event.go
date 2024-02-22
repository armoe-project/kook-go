package websocket

var sn = 0

type Payload struct {
	S  int         `json:"s"`
	D  interface{} `json:"d"`
	SN int         `json:"sn"`
}

func onPayload(payload *Payload) {
	if payload.S == 0 {
		sn = payload.SN
	}
}
