package main

import (
	"fmt"
)

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (NokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!!!!")
}

type IPhone struct {
}

func test() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
}
