package lib

import "fmt"

type MessageHandler func(msg, conn interface{}) interface{}

type Router struct {
	handlerMap map[int32]MessageHandler
}

func CreateRouter() *Router {
	router := &Router{}
	router.handlerMap = make(map[int32]MessageHandler)
	return router
}

func (router *Router) RegisterHandler(cmd int32, handler MessageHandler) {
	router.handlerMap[cmd] = handler
}

func (router *Router) Process(cmd int32, msg, conn interface{}) interface{} {
	handler, ok := router.handlerMap[cmd]
	if ok {
		res := handler(msg, conn)
		return res
	} else {
		fmt.Println("can't process cmd ", cmd)
	}
	return nil
}
