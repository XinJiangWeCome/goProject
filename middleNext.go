package main

import {
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	
}
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("request", "中间件")
		c.Next()
		c.writer
	}
}
