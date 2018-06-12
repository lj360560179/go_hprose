package main


import (
	"github.com/myzhan/boomer"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/satori/go.uuid"
)



func tcpc(){
	start := boomer.Now()
	client := rpc.NewTCPClient("tcp4://127.0.0.1:10001/")
	a ,_:= uuid.NewV1()
	client.Subscribe("OTHER", a.String(), nil, func(ip string) {
		elapsed := boomer.Now() - start
		if(ip == "SUCCESS"){
			boomer.Events.Publish("request_success", "tcp", "tcp", elapsed, int64(10))
		}
	})
}


func main(){
	task1 := &boomer.Task{
		Name: "tcpc",
		Weight: 10,
		Fn: tcpc,
	}

	boomer.Run(task1)

}