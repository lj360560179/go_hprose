package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
)
type event struct{}
func (e *event) OnError(name string, err error) {
	fmt.Printf("name: %s, err: %s\n", name, err.Error())
}

func main() {
	client := rpc.NewTCPClient("tcp4://127.0.0.1:8888/")
	client.SetEvent(&event{})
	count := 0
	done := make(chan struct{})
	client.Subscribe("time", "360560179", nil, func(data string) {
		count++
		if count > 10 {
			client.Unsubscribe("time")
			done <- struct{}{}
		}
		fmt.Println(data)
	})
	<-done
}