package main


import (
	"github.com/myzhan/boomer"
	"time"
)



func foo(){
	start := boomer.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := boomer.Now() - start
	boomer.Events.Publish("request_success", "http", "foo", elapsed, int64(10))
}

func bar(){
	start := boomer.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := boomer.Now() - start
	boomer.Events.Publish("request_failure", "udp", "bar", elapsed, "udp error")
}

func main(){
	task1 := &boomer.Task{
		Name: "foo",
		Weight: 10,
		Fn: foo,
	}
	task2 := &boomer.Task{
		Name: "bar",
		Weight: 20,
		Fn: bar,
	}
	boomer.Run(task1, task2)

}