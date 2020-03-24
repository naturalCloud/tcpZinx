package server

type Message struct {
	MessageId uint32
	DataLen   uint32
	Data      []byte
}

func (m *Message) GetMsgId() uint32 {
	return m.MessageId
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.MessageId = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
