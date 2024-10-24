package main

import (
	"fmt"
	"strconv"
)
func main(){
	var num int=42
	fmt.Printf("number is %d\n",num)
	fmt.Printf("Type of number is %T\n",num)
	var data float64=float64(num)
	data=data+1.23
	fmt.Println("data is ",data)
	fmt.Printf("Type of data is %T\n",data)
	a:=123
	str:=strconv.Itoa(a)
	fmt.Printf("str is %s, type is %T\n",str,str)
	number:="1234"
	b,_:=strconv.Atoi(number)
	b=b+1244
	fmt.Printf("b is %d, type is %T\n",b,b)
	str2:="3.14"
	c,_:=strconv.ParseFloat(str2,64)
	c=c+3.14
	fmt.Printf("c is %f, type is %T\n",c,c)
  



}