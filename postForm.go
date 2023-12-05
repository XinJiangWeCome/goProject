package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postForm() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		userpassword := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s, password: %s, type:%s", username, userpassword, types))
	})
	r.Run()
}
