package main

import "fmt"

func main() {
	if age := 20; age >= 18 {
		fmt.Println("age is greater than or equal to 18",age)
	}else{
		fmt.Println("age is less than 18",age)
	}
}