package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func goRedirect() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.51mh.com")
	})
	r.Run()
}
