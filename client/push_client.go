package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)
type PushService struct {
	Push func() ([]string, error)
	IdList func() ([]string,error)
}

func main() {

	var pushService *PushService
	client := rpc.NewClient("tcp://127.0.0.1:8888/")
	client.UseService(&pushService)

	router := gin.Default()
	router.GET("/push", func(context *gin.Context) {
		result ,_:=pushService.Push()
		sendResponse(result,context)
	})
	router.GET("/users", func(context *gin.Context) {
		result ,_:=pushService.IdList()
		sendResponse(result,context)
	})
	router.Run(":8000")
}
func sendResponse(data interface{},c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":data,
	})
}