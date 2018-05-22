package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"time"
)
func hello(name string, context *rpc.SocketContext) []string {
	ids :=make([]string,0)
	context.Clients().Broadcast("time", time.Now().String(), func(sended []string) {
		ids=append(ids,sended...)
	})
	return ids
}

func main() {
	server := rpc.NewTCPServer("tcp4://127.0.0.1:8888/")
	server.Publish("time", 0, 0)
	server.Event = event{}
	server.AddFunction("hello", hello)
	server.Start()
}

type event struct{}

func (event) OnSubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("客户 " + id + " 订阅了 topic: " + topic)
}
func (event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("客户 " + id + " 离线: " + topic)
}