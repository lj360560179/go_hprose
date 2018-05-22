package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)
type HelloService struct {
	Hello func() ([]string, error)
}

func main() {

	var helloService *HelloService
	client := rpc.NewClient("tcp://127.0.0.1:8888/")
	client.UseService(&helloService)


	router := gin.Default()
	router.GET("/push", func(context *gin.Context) {
		result ,_:=helloService.Hello()
		sendResponse(result,context)
	})
	router.GET("/users", func(context *gin.Context) {

	})
	router.Run(":8000")




}
func sendResponse(data interface{},c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":data,
	})
}