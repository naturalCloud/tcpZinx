package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"netLearn/netlib/sInterface"
	"netLearn/netlib/util"
)

type DataPack struct {
}

const DataHeadLen = 8 //定义数据head头

//获取数据包头的默认长度
func (d *DataPack) GetHeadLen() uint32 {
	//4字节长度 unit32  id + 4字节 data len uint32
	return DataHeadLen
}

//打包数据为二级制
func (d *DataPack) Pack(msg sInterface.Message) ([]byte, error) {
	//创建存放字节的缓冲
	buffer := bytes.NewBuffer([]byte{})

	//写入head 包头数据 数据长度
	err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}

	//写入head 包数据 消息id
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写入data

	if err := binary.Write(buffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

//解包数据
func (d *DataPack) UnPack(data []byte) (sInterface.Message, error) {
	databuf := bytes.NewReader(data)

	msg := &Message{}

	//先读数据长度 4字节
	if err := binary.Read(databuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读数据id
	if err := binary.Read(databuf, binary.LittleEndian, &msg.MessageId); err != nil {
		return nil, err
	}

	//读数据
	if util.ServerConf.MaxBufSize > 0 && msg.DataLen > util.ServerConf.MaxBufSize {
		return nil, errors.New("data too large ")
	}
	return msg, nil
}

func NewDataPack() *DataPack {
	return &DataPack{}

}
