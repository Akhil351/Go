package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums==nil)
	fmt.Println(nums)
	fmt.Println(len(nums))
	var a=make([]int,0,5)
	a=append(a, 1,2,3,4,5,6)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	b:=[]int{}
	b=append(b, 1,2)
	fmt.Println(b)
	fmt.Println(len(b),cap(b))
	var c=make([]int,0,5)
	c=append(c, 1,2,3)
	var d=make([]int,len(c))
	copy(d,c)
	fmt.Println(c,d)
	fmt.Println(d[0:2])
	fmt.Println(d[1:])
	fmt.Println(d[:3])
	d[0]=9
	fmt.Println(slices.Equal(c,d))
	 mat:=[][]int{{1,2,3},{1,2,3}}
	fmt.Println(mat)
}