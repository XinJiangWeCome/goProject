package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginUrl() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run()
}
