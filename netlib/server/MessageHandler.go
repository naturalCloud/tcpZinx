package server

import (
	"fmt"
	"netLearn/netlib/sInterface"
)

type MessageHandler struct {
	MsgHandleMap map[uint32]sInterface.Router
}


//添加路由到map集合 key 为msgId ,value 为 Router
func (m *MessageHandler) AddRouterMap(msgId uint32, router sInterface.Router) {
	if _, ok := m.MsgHandleMap[msgId]; ok {
		fmt.Println("current router exits", msgId)
		return
	}
	m.MsgHandleMap[msgId] = router
	fmt.Println("add router to routerMap success",msgId)
}

func (m *MessageHandler) DoMessageHandle(request sInterface.Request) {
	handler ,ok := m.MsgHandleMap[request.GetMsgId()]
	if !ok {
		fmt.Println("api router not found ,must reg")
		return
	}
	//处理消息路由
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PreHandle(request)
}

//初始化 message
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		MsgHandleMap: make(map[uint32]sInterface.Router),
	}
}
