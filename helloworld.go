package main

import (
	"fmt"
)

var g int

func caseIt() {
	// %d 表示整型数字，%s 表示字符串
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a 变量得地址： %x\n", &a)
	fmt.Printf("ip 变量存储得指针地址：%x\n", ip)
	fmt.Printf("ip 变量的值： %d\n", *ip)
}
