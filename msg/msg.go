package msg

import "github.com/zeromicro/go-zero/core/jsonx"

type D map[string]interface{}

type Msg struct {
	Id        string `json:"id"`
	Timestamp uint64 `json:"timestamp"`
	Type      string `json:"type"`
	Data      D      `json:"data"`
}

// ToBytes Msg to uint8[]
func (m *Msg) ToBytes() []byte {
	marshal, _ := jsonx.Marshal(m)
	return marshal
}

// ToMsg uint8[] to Msg
func ToMsg(data []byte) *Msg {
	msg := &Msg{}
	_ = jsonx.Unmarshal(data, msg)
	return msg
}
