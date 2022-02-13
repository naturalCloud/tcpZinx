package server

import (
	"errors"
	"fmt"
	"netLearn/netlib/sInterface"
	"sync"
)

type ConnectionManager struct {
	connectionsMap map[uint32]sInterface.Connection

	connLock sync.RWMutex
}

func (c *ConnectionManager) AddConn(conn sInterface.Connection) {

	//共享读写锁map
	c.connLock.Lock()
	defer c.connLock.Unlock()

	//添加链接到map
	c.connectionsMap[conn.GetConnId()] = conn

	fmt.Printf("connId %d add to connMap succ , current map len --> %d \n", conn.GetConnId(), c.Len())

}

//删除链接
func (c *ConnectionManager) RemoveConn(conn sInterface.Connection) {
	//添加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connectionsMap[conn.GetConnId()]; ok {
		delete(c.connectionsMap, conn.GetConnId())
		fmt.Printf("connId %d remove to connMap succ, connMap len -- %d \n", conn.GetConnId(), c.Len())
		return
	}

	fmt.Printf("connId %d not found \n",conn.GetConnId())

}

//获取当前链接
func (c *ConnectionManager) GetConn(connId uint32) (sInterface.Connection, error) {
	//添加读锁getMap
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connectionsMap[connId]; ok {
		fmt.Printf("connId %d found \n", conn.GetConnId())
		return conn, nil
	}
	return nil, errors.New(fmt.Sprintf("connId %d not found", connId))
}

//获取链接数
func (c *ConnectionManager) Len() int {
	return len(c.connectionsMap)
}

//清理链接
func (c *ConnectionManager) ClearConn() {
	//添加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connId, conn := range c.connectionsMap {
		//停止链接
		conn.Stop()
		//移除链接
		delete(c.connectionsMap, connId)
	}

	fmt.Println("connMap clear ok,connNum", c.Len())

}

func NewConnectionManager() *ConnectionManager {

	return &ConnectionManager{
		connectionsMap: make(map[uint32]sInterface.Connection),
		connLock:       sync.RWMutex{},
	}
}
