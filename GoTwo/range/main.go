package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	sum:=0
	index:=0
	for i,num:=range nums{
		sum+=num
		index+=i
	}
	fmt.Println(sum,index)

	m:=map[string]string{"name":"golang","area":"backend"}
	for key,value:=range m{
		fmt.Println(key,value)
	}
	 for i,c:=range "goLang"{
		//fmt.Printf("%d %c\n",i,c)
		fmt.Println(i,string(c))
	 }
}