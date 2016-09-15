package main

import (
	"fmt"
)

func main() {
	//5: 函数作为方法
	var c Circle
	c.radius = 10.00
	area := c.getArea()
	fmt.Println("area: ", area)
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

