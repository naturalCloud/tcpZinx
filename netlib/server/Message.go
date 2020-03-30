package server

type Message struct {
	MessageId uint32
	DataLen   uint32
	Data      []byte
}

func (m *Message) GetMsgId() uint32 {
	return m.MessageId
}

func (m *Message) GetMsgLen() uint32 {
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

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		MessageId: id,
		DataLen:   uint32(len(data)),
		Data:      data,
	}

}

func BuildTextMsg(msgId uint32, data string) *Message {
	return NewMessage(msgId, []byte(data))
}
