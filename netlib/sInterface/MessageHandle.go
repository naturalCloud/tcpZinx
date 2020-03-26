package sInterface

type MessageHandle interface {
	AddRouterMap(uint32, Router)
	DoMessageHandle(Request)
}
