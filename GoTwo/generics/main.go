package main

import "fmt"
// [T any]
// [T int | string]
func printSlice[T comparable,V any](items []T,name V) {
	for _, item := range items {
		fmt.Print(item, " ",name," ")
	}
	fmt.Println()
}

type Stack[T any] struct{
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4}
	printSlice(nums,"John")
	stack:=Stack[string]{
		elements: [] string {"golang","java"},
	}
	fmt.Println(stack,"John")
}
