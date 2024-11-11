package middleware

import (
	"fmt"
	"gin-django-example/utils"
	"github.com/gin-gonic/gin"
)

// LogMiddleware 日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTimeStr := utils.GetNowTimeFormat("2006-01-02 15:04:05.000")
		fmt.Println("接收time:%v", nowTimeStr)
		method := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		ip := c.ClientIP()
		params := ""
		for k, v := range c.Request.PostForm {
			params += fmt.Sprintf("%s:%v ", k, v)
		}
		fmt.Println(fmt.Sprintf("method:%v;reqUrl:%v;statusCode:%v;ip:%v;params:%v", method, reqUrl, statusCode, ip, params))
		c.Next()
		nowTimeStr2 := utils.GetNowTimeFormat("2006-01-02 15:04:05.000")
		fmt.Println("返回time:%v", nowTimeStr2)
	}
}
