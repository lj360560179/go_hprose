package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"fmt"
)
type event struct {}

func (e *event) OnError(name string, err error) {
	fmt.Printf("name: %s, err: %s\n", name, err.Error())
}

func main() {
	client := rpc.NewTCPClient("tcp4://127.0.0.1:8888/")
	client.SetEvent(&event{})
	done := make(chan struct{})
	client.Subscribe("push", "704c95471f834d3488c76bb2b9bd63b61-AND", nil, func(ip string) {
		done <- struct{}{}
		fmt.Println(ip)
	})
	<-done
}
