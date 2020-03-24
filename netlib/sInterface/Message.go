package sInterface

type Message interface {
	GetMsgId() uint32   //获取消息id
	GetDataLen() uint32 //获取消息长度
	GetData() []byte    //获取消息内容

	SetMsgId(uint32)   //设置消息id
	SetData([]byte)    //设置消息内容
	SetDataLen(uint32) //设置消息长度
}
