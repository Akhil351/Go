package main

import "fmt"

func fun1(num *int) {
	*num = *num * 2
}
func main() {
	//var num int
	num := "user"
	//var ptr *int
	ptr := &num
	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}
	value := 10
	fun1(&value)
	fmt.Println(value)
}
