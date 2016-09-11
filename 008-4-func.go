package main

import (
	"fmt"
)

func main() {
	//4, 函数闭包
	c := getSequence()
	fmt.Println("c(): ", c())
	fmt.Println("c(): ", c())
	fmt.Println("c(): ", c())

	c2 := getSequence()
	fmt.Println("c2(): ", c2())
	fmt.Println("c2(): ", c2())
	fmt.Println("c2(): ", c2())
}

//4, 函数闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}
