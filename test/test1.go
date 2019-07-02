package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 指针。指针保存了值的内存地址。即间接引用和重定向
	var p *int // 声明p的类型是指向int类型的指针，零值为nil

	var i = 42

	p = &i //& 操作符会生成一个指向其操作数的指针。

	fmt.Println(*p)

	//fmt.Println(Books{"Go 语言", "www.baidu.com", "haha", 1})
	//
	//fmt.Println(Books{title: "PHP", author: "www.baidu.com", subject: "lueluelue", book_id: 2})
	//
	//// 不传的话默认为空或者0
	//fmt.Println(Books{title: "php"})

	book := Books{"Go 语言", "www.baidu.com", "haha", 1}
	fmt.Println(book.title)

	// 普通赋值更改book1 不会修改book原本的值
	book1 := book
	book1.title = "php"
	fmt.Println(book1.title, "\n", book.title)

	// 用指针的话会修改原来的值
	book2 := &book
	book2.title = "Java"
	fmt.Println(book2.title, "\n", book.title)
}
