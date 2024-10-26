package main

import "fmt"

func main() {
	var nums [4]int
	nums[0]=1
	nums[1]=2
	var vals [4] bool
	vals[2]=true
	var names [4] string
	names[1]="goLang"
	fmt.Println(len(nums))
	fmt.Println(nums)
	fmt.Println(vals)
	fmt.Println(names)
	a:=[3]int{1,3,4}
	fmt.Println(a)
	b:=[2][2]int{{1,2},{3,4}}
	fmt.Println(b) 
}