package main

import (
	"fmt"
)

func main(){
	//声明并初始化一个数组
	arr := [] int {1, 2, 3, 4}
	//遍历数组
	for i := 0; i<4; i++{
		fmt.Printf("i: %d\n", arr[i])
	}

	//数组作为函数参数
	avg := getAverage(arr, 4)
	fmt.Printf("avg: %f\n", avg)

}

func getAverage(arr []int, len int) float32{
	sum := 0
	for i:=0; i< len; i++{
		sum += arr[i]
	}

	return float32(sum/len)
}

