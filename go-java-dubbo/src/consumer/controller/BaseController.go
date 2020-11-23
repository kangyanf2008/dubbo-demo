package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-java-dubbo/api-model/model"
	"go-java-dubbo/consumer/service"
	"net/http"
	"time"
)

var (
	server             *http.Server
)


//启动web服务
func Run() {
	router := gin.New()
	router.GET("/consumer/index/:table", index)
	server = &http.Server{
		Handler: router,
	}
	router.Run("0.0.0.0:81") // listen and serve on 0.0.0.0:port

}

func index(c *gin.Context) {
	table := c.Param("table")
	b := &model.RequestData{Table: table}
	//responseData, _:= service.DataService.DataList(b)
	responseData := &model.ResponseData{}
	err := service.DataService.DataList(context.Background(), b, responseData)
	fmt.Println(err)
	c.String(200, fmt.Sprintf("Hello World #%v!",responseData))
}


//服务退出
func Shutdown() {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}
}

