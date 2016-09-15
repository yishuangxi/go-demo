package main

import (
	"fmt"
)

func main(){
	//指针作为函数参数
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a的值: %d\n", a)
	fmt.Printf("交换前 b的值: %d\n", b)

	swapPtr(&a, &b)

	fmt.Printf("交换后 a的值: %d\n", a)
	fmt.Printf("交换后 b的值: %d\n", b)
}

func swapPtr(a, b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

