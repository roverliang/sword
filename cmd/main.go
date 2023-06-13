package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roverliang/sword/logger"
	"github.com/roverliang/sword/test"
	"net/http"
)

func init() {
	logger.Init()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	logger.Infof("test: %s", "hello world")
	logger.Errorf("faild : %s", "jkdjflsdfja")
	test.AddLog()
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
