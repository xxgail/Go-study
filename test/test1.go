package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	//fmt.Println(Books{"Go 语言", "www.baidu.com", "haha", 1})
	//
	//fmt.Println(Books{title: "PHP", author: "www.baidu.com", subject: "lueluelue", book_id: 2})
	//
	//// 不传的话默认为空或者0
	//fmt.Println(Books{title: "php"})

	book := Books{"Go 语言", "www.baidu.com", "haha", 1}
	fmt.Println(book.title)

	book1 := book
	book1.title = "php"
	fmt.Print(book1.title, book.title)
}
