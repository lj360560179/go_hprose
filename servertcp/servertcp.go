package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strings"
)


func main() {

	server := rpc.NewTCPServer("tcp4://127.0.0.1:8888/")
	server.Publish("OTHER1", 10000, 0)
	server.Event = event{}
	//开协程监听tcp
	go server.Start()

	//http路由
	router := gin.Default()
	//发送推送消息
	router.GET("/push", func(context *gin.Context) {
		server.Broadcast("OTHER", time.Now().String(), func(sended []string) {
			fmt.Println(sended)
		})
	})
	//获取在线用户
	router.GET("/idList", func(context *gin.Context) {
		users := server.IDList("OTHER")
		result := make([]string ,0)
		for _,v:= range users{
			result = append(result,strings.Split(v,"-")[0])
		}
		sendResponse(result,context)
	})
    router.Run("127.0.0.1:8000")
}

type event struct{}

func (event) OnSubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("用户" + id + " 上线: " + topic)
}
func (event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("用户" + id + " 离线: " + topic)
}

func sendResponse(data interface{},c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":data,
	})
}