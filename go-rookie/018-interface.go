package main

import (
	"fmt"
)

func main(){
	var phone1 Phone
	var phone2 Phone
	phone1 = new(NokiaPhone)
	phone2 = new(IPhone)

	phone1.call()
	phone2.call()
}

type Phone interface {
	call()
}

type IPhone struct {

}

type NokiaPhone struct {
}

func (iPhone *IPhone) call(){
	fmt.Println("iphone calling")
}

func (nokiaPhone *NokiaPhone) call(){
	fmt.Println("nokiaphone calling")
}



