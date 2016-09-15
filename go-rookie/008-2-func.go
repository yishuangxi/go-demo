package main

import (
	"fmt"
)

func main() {

	//2, 函数返回多个值
	a, b := swap(2, 3)
	fmt.Println("x, y : ", a, b)
}

func swap(x, y int) (int, int) {
	return y, x
}

