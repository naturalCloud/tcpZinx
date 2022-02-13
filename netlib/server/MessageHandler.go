package server

import (
	"fmt"
	"netLearn/netlib/sInterface"
	"netLearn/netlib/util"
)

type MessageHandler struct {
	MsgHandleMap map[uint32]sInterface.Router //路由map存放messageId => router

	WorkPoolSize       uint32                    //worker 协程池的数量
	TaskWorkQueue      []chan sInterface.Request //取worker投递的消息chain
	workPoolInitialize bool                      //work 池是否初始化
}

func (m *MessageHandler) WorkPoolIsInit() bool {
	return m.workPoolInitialize
}

//StartWorkerPool() 方法是启动Worker工作池， 这里根据用户配置好
//的 WorkerPoolSize 的数量来启动， 然后分别给每个Worker分配一
//个 TaskQueue ， 然后用一个goroutine来承载一个Worker的工作业务。
func (m *MessageHandler) StarWorkPool() {

	if m.WorkPoolSize < 1 {
		panic("work poolSize < 1 , stop init work pool " )
		return
	}
	for i := 0; i < int(m.WorkPoolSize); i++ {
		m.TaskWorkQueue[i] = make(chan sInterface.Request, util.ServerConf.MaxTaskQueueLen)
		//启动一个work 协程
		go m.startOneWork(i, m.TaskWorkQueue[i])
	}
	m.workPoolInitialize = true
	fmt.Printf("work pool init success ,WorkPoolSize is %d \n", m.WorkPoolSize)
}

//启动一个work 进行任务处理
//StartOneWorker() 方法就是一个Worker的工作业务， 每个worker是不会退出的
//(目前没有设定worker的停止工作机制)， 会永久的从对应的TaskQueue中等待消
//息， 并处理
func (m *MessageHandler) startOneWork(workId int, taskQueue chan sInterface.Request) {

	fmt.Printf("this is %d worker runing  \n", workId)
	for {
		select {
		case request := <-taskQueue:
			m.DoMessageHandle(request)
			fmt.Printf("workId [%d ] handel over , current link requestId [%d] \n ", workId, request.GetRequestId())
		}

	}

}

//将消息投递给某个worker进行处理
func (m *MessageHandler) SendMsgToTaskQueue(request sInterface.Request) {

	//StartOneWorker() 方法就是一个Worker的工作业务， 每个worker是不会退出的
	//(目前没有设定worker的停止工作机制)， 会永久的从对应的TaskQueue中等待消
	//息， 并处理
	workId := request.GetRequestId() % m.WorkPoolSize

	fmt.Printf(" ConnId [%d]  link data Send to [%d] worker handel ", request.GetConn().GetConnId(), workId)
	m.TaskWorkQueue[int(workId)] <- request

	//SendMsgToTaskQueue() 作为工作池的数据入口， 这里面采用的是轮询的分配机
	//制， 因为不同链接信息都会调用这个入口， 那么到底应该由哪个worker处理该链接
	//的请求处理， 整理用的是一个简单的求模运算。 用余数和workerID的匹配来进行分
	//配。

}

//添加路由到map集合 key 为msgId ,value 为 Router
func (m *MessageHandler) AddRouterMap(msgId uint32, router sInterface.Router) {
	if _, ok := m.MsgHandleMap[msgId]; ok {
		fmt.Println("current router exits", msgId)
		return
	}
	m.MsgHandleMap[msgId] = router
	fmt.Println("add router to routerMap success", msgId)
}

func (m *MessageHandler) DoMessageHandle(request sInterface.Request) {
	handler, ok := m.MsgHandleMap[request.GetMsgId()]
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
		MsgHandleMap:       make(map[uint32]sInterface.Router),
		WorkPoolSize:       util.ServerConf.WorkPoolSize,
		TaskWorkQueue:      make([]chan sInterface.Request, util.ServerConf.MaxTaskQueueLen),
		workPoolInitialize: false,
	}
}
