package main

import (
	"fmt"
)

func main(){
	const PI float32 = 3.14
	//以下代码, 如果重新赋值会报错
	// PI = 3.1415926
	fmt.Println("PI: ", PI)
}

