package main

import (
	"fmt"
)

func main(){
	//声明一个整型变量,并赋值
	var a int = 20
	//声明一个整型指针, 并赋值
	var ip *int = &a


	fmt.Printf("a变量的地址是: %x\n", &a)
	fmt.Printf("ip的值是: %x\n", ip)
	fmt.Printf("使用指针访问变量a的值, *ip: %d\n", *ip)

	var ptr *int
	if ptr == nil{
		fmt.Printf("ptr是空指针, 值为%x\n", ptr)
	}else{
		fmt.Printf("ptr不是空指针, 值为%s\n", ptr)
	}

}

