package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", GetData)
	r.GET("/detail/:id", GetDetail)
	// r.Run(":9090")
	fmt.Println("gin启动成功")
	return r
}
