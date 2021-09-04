package client

import "spike.io/bin"

type Topic struct {
	Topic   string `json:"topic,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	Offset  int64  `json:"offset,omitempty"`
}

func (m *Topic) ProtoMessage() *bin.Topic {
	return &bin.Topic{
		Topic:   m.Topic,
		GroupId: m.GroupId,
		Offset:  m.Offset,
	}
}

type Success struct {
	Success bool   `json:"success,omitempty"`
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func protoToSuccess(message *bin.Success) *Success {
	return &Success{
		Success: message.GetSuccess(),
		Code:    message.GetCode(),
		Message: message.GetMessage(),
	}
}

type Message struct {
	Topic   string            `json:"topic,omitempty"`
	GroupId string            `json:"group_id,omitempty"`
	Offset  int64             `json:"offset,omitempty"`
	Value   []byte            `json:"value,omitempty"`
	Header  map[string][]byte `json:"header,omitempty"`
}

func protoToMessage(message *bin.Message) *Message {
	return &Message{
		Topic:   message.GetTopic(),
		GroupId: message.GetGroupId(),
		Offset:  message.GetOffset(),
		Value:   message.GetValue(),
		Header:  message.GetHeader(),
	}
}

func (m *Message) ProtoMessage() *bin.Message {
	return &bin.Message{
		Topic:   m.Topic,
		GroupId: m.GroupId,
		Offset:  m.Offset,
		Value:   m.Value,
		Header:  m.Header,
	}
}
