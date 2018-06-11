package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"os"
)
type event struct {}

func (e *event) OnError(name string, err error) {
	fmt.Printf("name: %s, err: %s\n", name, err.Error())
}

var result = make(chan string, 1)

func main() {
	for i:=0;i<10000;i++ {
		go func() {
			a ,_:= uuid.NewV1()
			client := rpc.NewTCPClient("tcp4://172.16.1.102:10001/")
			//client := rpc.NewTCPClient("tcp4://127.0.0.1:8888/")
			client.SetEvent(&event{})
			client.Subscribe("OTHER", a.String(), nil, func(ip string) {
				result<-ip
			})
		}()
	}

	fileName := "./a.log"
	dstFile,err := os.Create(fileName)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	for v:= range result{
		if v != "SUCCESS"{
			dstFile.WriteString(v+"\n")
		}
	}

	time.Sleep(time.Second*3000000)
}
