package main

import (
	"fmt"
)

func main() {
	//结构体指针
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

	//定义结构体指针
	var ptr_book1 *Book
	var ptr_book2 *Book

	//使用&符号,从变量Book1和Book2中获取变量地址,并赋值给指针变量
	ptr_book1 = &Book1
	ptr_book2 = &Book2

	printBookInfoByPtr(ptr_book1)
	fmt.Printf("\n")
	printBookInfoByPtr(ptr_book2)

}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBookInfoByPtr(book *Book) {
	//对于结构体类型,访问值不需要加*, 比如: *book.title
	fmt.Printf("Book title : %s\n", book.title);
	fmt.Printf("Book author : %s\n", book.author);
	fmt.Printf("Book subject : %s\n", book.subject);
	fmt.Printf("Book book_id : %d\n", book.book_id);
}


