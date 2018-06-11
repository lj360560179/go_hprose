package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
)
type event struct {}

func (e *event) OnError(name string, err error) {
	fmt.Printf("name: %s, err: %s\n", name, err.Error())
}

func main() {
	count:=0
	for i:=0;i<20000;i++ {
		go func() {
			a ,_:= uuid.NewV1()

			client := rpc.NewTCPClient("tcp4://172.16.1.102:10001/")
			//client := rpc.NewTCPClient("tcp4://127.0.0.1:8888/")
			client.SetEvent(&event{})
			client.Subscribe("OTHER", a.String(), nil, func(ip string) {
				count++
				fmt.Println(ip)
			})
		}()
	}

	time.Sleep(time.Second*3000000)
}
