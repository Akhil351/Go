package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(time.Second * 2)
	fmt.Println("hello function completed")
}
func sayHi() {
	fmt.Println("Hi")
	time.Sleep(time.Second)
	fmt.Println("hi function completed")
}

func main() {
	fmt.Println("main function ")
	go sayHello()
	go sayHi()
	time.Sleep(time.Second * 3)
}
