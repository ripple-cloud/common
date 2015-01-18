package message

// TODO: Add tests for this package

import (
	"encoding/json"
	"io"
)

type MsgType int

type Message struct {
	Type MsgType           `json:"type"`
	Meta map[string]string `json:"meta",omitempty`
	Body []byte            `json:"body",omitempty`
}

const (
	Noop MsgType = iota

	// messages a service will send
	Register
	Deregister
	Publish

	// messages a service will receive
	Request
	Ack
	Error
)

func (mt MsgType) String() string {
	switch mt {
	case Register:
		return "register"
	case Deregister:
		return "deregister"
	case Publish:
		return "publish"
	case Request:
		return "request"
	case Ack:
		return "ack"
	case Error:
		return "error"
	}
	return ""
}

func New() *Message {
	return &Message{
		Type: Noop,
		Meta: map[string]string{},
		Body: []byte{},
	}
}

func NewRegister(topic string, meta map[string]string) *Message {
	meta["topic"] = topic
	return &Message{
		Type: Register,
		Meta: meta,
		Body: []byte{},
	}
}

func NewDeregister(topic string, meta map[string]string) *Message {
	meta["topic"] = topic
	return &Message{
		Type: Deregister,
		Meta: meta,
		Body: []byte{},
	}
}

func NewRequest(topic string, meta map[string]string, body []byte) *Message {
	meta["topic"] = topic
	return &Message{
		Type: Request,
		Meta: meta,
		Body: body,
	}
}

func NewPublish(topic string, meta map[string]string, body []byte) *Message {
	meta["topic"] = topic
	return &Message{
		Type: Publish,
		Meta: meta,
		Body: body,
	}
}

func Decode(r io.Reader) (*Message, error) {
	m := &Message{}
	err := json.NewDecoder(r).Decode(m)
	return m, err
}

func (m *Message) Encode() ([]byte, error) {
	return json.Marshal(m)
}
