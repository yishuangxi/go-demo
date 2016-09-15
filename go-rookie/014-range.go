package main

import (
	"fmt"
)

func main() {

	//1, 枚举数组元素
	var arr []int = []int{10, 20, 30, 40}
	for index, value := range arr {
		fmt.Printf("index: %d value: %d\n", index, value)
	}

	//2, 枚举字符串中的字符
	var s string = "abcdefg"
	for index, char := range s {
		fmt.Printf("index: %d char: %d\n", index, char)
	}
}

