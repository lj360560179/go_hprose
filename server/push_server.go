package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"time"
)
func push(context *rpc.SocketContext) []string {
	ids :=make([]string,10)
	context.Clients().Broadcast("push", time.Now().String(), func(sended []string) {
		ids=append(ids,sended...)
	})
	return ids
}

func idList(context *rpc.SocketContext)  []string{
	return context.Clients().IDList("push")
}

func main() {
	server := rpc.NewTCPServer("tcp4://127.0.0.1:8888/")
	server.Publish("time", 0, 0)
	server.Event = event{}
	server.AddFunction("push", push)
	server.AddFunction("idList", idList)
	server.Start()
}

type event struct{}

func (event) OnSubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("用户" + id + " 上线: " + topic)
}
func (event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("用户" + id + " 离线: " + topic)
}