package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func ()  {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <-i
			fmt.Println("子go程正在运行： len = ", len(c), ", 元素=", i, ", cap = ", cap(c))
		}

	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("end...")

}