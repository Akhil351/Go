package main

import "fmt"

// in go only forLoop is used for looping
func main(){
	i:=1
    for i<=3 {
		fmt.Println(i)
		i=i+1
	}
	// infinte loop
	// for{
	// 	fmt.Println(1)
	// }
	// classic for loop
	for i:=0;i<3;i++{
		if i==1{
			continue
		}
		fmt.Println(i)
	}
	// using range
	fmt.Println("range")
	for j:= range 11{
		fmt.Println(j)
	}
}