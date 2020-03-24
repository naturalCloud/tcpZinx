package sInterface

type DataPack interface {
	//获取包头长度
	GetHeadLen() uint32
	//封装包体
	Pack(Message) ([]byte, error)
	//解包
	UnPack([]byte) (Message, error)
}
