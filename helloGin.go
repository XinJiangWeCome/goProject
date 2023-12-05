package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloGin() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!!!")
	})
	r.Run(":8080")
}
