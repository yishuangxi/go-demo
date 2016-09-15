package main

import (
	"fmt"
)

func main() {
	//结构体定义
	var Book1 Book
	var Book2 Book
	
	Book1.title = "Go语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go语言教程"
	Book1.book_id = 1
	
	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python语言教程"
	Book2.book_id = 2
	
	fmt.Printf("Book1 title: %s\n", Book1.title)
	fmt.Printf("Book1 author: %s\n", Book1.author)
	fmt.Printf("Book1 subject: %s\n", Book1.subject)
	fmt.Printf("Book1 book_id: %d\n", Book1.book_id)

	fmt.Printf("\n")

	fmt.Printf("Book2 title: %s\n", Book2.title)
	fmt.Printf("Book2 author: %s\n", Book2.author)
	fmt.Printf("Book2 subject: %s\n", Book2.subject)
	fmt.Printf("Book2 book_id: %d\n", Book2.book_id)
}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}


