package main

import "fmt"

func main() {
	var num int = 100
	var ptr *int
	ptr = &num
	fmt.Println("ptr 指针的值：", ptr)
	fmt.Println("*ptr指针的值：", *ptr)
	num = 200
	fmt.Println("*ptr指针的值：", *ptr)
}
