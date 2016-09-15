package main

import (
	"fmt"
)

func main() {

	//1, 函数基本使用方式
	max := maxNumber(100, 2)
	fmt.Println("max: ", max)
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

