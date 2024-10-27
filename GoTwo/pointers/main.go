package main

import "fmt"
func change(a *int){
	*a=5
	fmt.Println("in change function",*a)
}
func main() {
	num := 1
	change(&num)
	var a,b int=2,5
	ptr1:=&a
	ptr2:=&b
	sum:=*ptr1+*ptr2
	fmt.Println("in main function",num)
	fmt.Println(sum)
}
