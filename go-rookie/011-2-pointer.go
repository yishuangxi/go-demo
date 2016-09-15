package main

import (
	"fmt"
)

const MAX int = 3

func main() {
	//指针数组
	var a []int = []int{10, 20, 30}

	var ptr [MAX]*int

	//遍历,将数组a里面的每一个值的地址赋值给指针数组的每一个元素
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("ptr[%d] = %x\n", i, ptr[i])
	}

	//通过指针数组来访问原数组的值
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

