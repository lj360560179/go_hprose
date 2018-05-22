package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)
type HelloService struct {
	Hello func(string) (string, error)
}

func main() {

	router := gin.Default()
	router.GET("/push", func(context *gin.Context) {

	})
	router.GET("/users", func(context *gin.Context) {

	})
	router.Run(":8000")


	client := rpc.NewClient("tcp://127.0.0.1:8888/")
	client.Subscribe("ip", "", nil, func(ip string) {
		fmt.Println(ip)
	})
	var helloService *HelloService
	client.UseService(&helloService)
	for i := 0; i < 10; i++ {
		fmt.Println(helloService.Hello("world"))
	}
}
func sendResponse(data interface{},c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":data,
	})
}