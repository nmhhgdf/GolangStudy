package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	b := &Book{}

	var r Reader
	r = b

	// 查看 r 的类型
	fmt.Printf("r 的类型: %T\n", r) // 输出: *main.Book
	fmt.Printf("r 的值: %v\n", r)  // 输出: 地址值

	// 类型断言
	if w, ok := r.(Writer); ok {
		fmt.Println("\n类型断言成功!")
		fmt.Printf("w 的类型: %T\n", w)              // 输出: *main.Book
		fmt.Printf("w 和 b 是同一个对象吗? %v\n", w == b) // 输出: true

		w.WriteBook() // 调用 WriteBook
	}

	// 也可以使用类型 switch
	switch v := r.(type) {
	case Writer:
		fmt.Println("\n类型 switch 也确认实现了 Writer 接口")
		v.WriteBook()
	default:
		fmt.Println("不是 Writer")
	}
}
