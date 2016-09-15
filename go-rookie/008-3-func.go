package main

import (
	"fmt"
)

func main() {

	//3, 函数参数传递指针
	x, y := 1, 2
	fmt.Println("before swap: ", x, y)
	swapByPtr(&x, &y)
	fmt.Println("after: swap: ", x, y)
}

func swapByPtr(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

