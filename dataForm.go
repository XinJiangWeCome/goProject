package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func dataForm() {
	r := gin.Default()
	// JSON响应
	r.GET("/soneJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "soneJSON", "status": 200})
	})
	// 结构体相应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Message = "message"
		msg.Name = "root"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// XML
	r.GET("/someXML", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{"message": "abc"})
	})
	// YAML
	r.GET("/someYAML", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "zhangsan"})
	})
	// protobuf格式
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run(":8000")
}
