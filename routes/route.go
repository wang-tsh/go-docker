package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-docker/config"
	"go-docker/handler"
)

func Api() *gin.Engine {
	api := gin.Default()
	api.Use(ReqLog())
	makeApi(api)
	return api
}

func makeApi(api *gin.Engine) {
	mainGroup := api.Group(config.AppConfig.ContextPath)
	makeBusApis(mainGroup)
}

func makeBusApis(group *gin.RouterGroup) {
	accessSystemGroup := group.Group("/accessSystem")
	makeAccessSystemApis(accessSystemGroup)
}

func makeAccessSystemApis(group *gin.RouterGroup) {
	group.GET("/systems", handler.GetSystems)
}

func ReqLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求参数
		get, _ := c.Get("req")
		marshal, _ := json.Marshal(get)
		// 打印日志
		logrus.WithFields(logrus.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqUrl,
			"body":        string(marshal),
		}).Info()
	}
}
