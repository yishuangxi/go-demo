package main

import (
	"fmt"
)

func main() {
	//结构体作为函数参数
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

	printBookInfo(Book1)
	fmt.Printf("\n")
	printBookInfo(Book2)

}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBookInfo(book Book) {
	fmt.Printf("Book title : %s\n", book.title);
	fmt.Printf("Book author : %s\n", book.author);
	fmt.Printf("Book subject : %s\n", book.subject);
	fmt.Printf("Book book_id : %d\n", book.book_id);
}


