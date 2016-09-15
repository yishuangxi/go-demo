package main

import (
	"fmt"
)

func main() {
	//bool类型
	var tr = true
	var fa = false
	if tr {
		fmt.Println("这里是bool true")
	}
	if fa {
		fmt.Println("这里是bool false, 永远不会被打印出来")
	}

	//数字类型
	var i int = 100
	var f float64 = 12.123456
	fmt.Println("i是一个int型数据, 值为:", i)
	fmt.Println("f是一个float类型数据, 值为: ", f)
}
