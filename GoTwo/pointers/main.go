package main

import "fmt"
func change(num *int){
	*num=5
	fmt.Println("in change function",*num)
}
func main() {
	num := 1
	change(&num)
	fmt.Println("in main function",num)
}
