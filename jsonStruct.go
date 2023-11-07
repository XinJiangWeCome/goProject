package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	r := gin.Default()
	b := gin.H{"status": "登录成功"} //GIN.H就是将数据转换成map类型
	// json的转换
	var u user
	u = user{101, "lisi", 28}
	userData, err := json.Marshal(u) //json与对象之间转换,err记录的是错误信息，只有当错误信息为空的时候方可打印数据 unMarshal是json转换成对象
	if err != nil {
		fmt.Print("error in it")
	} else {
		fmt.Print(string(userData))
	}
	for key, value := range b {
		fmt.Printf(key, value)
	}
	r.Run(":8082")
}
