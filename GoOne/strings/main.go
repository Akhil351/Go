package main

import (
	"fmt"
	"strings"
)
func main(){
	data:="apple.orange.banana.pineapple"
	parts:=strings.Split(data,".")
	fmt.Println(parts)
	str:="one two three two two five two"
	count:=strings.Count(str,"two")
	fmt.Println(count)
	a:="  Hello Venom   "
	trimmed:=strings.TrimSpace(a)
	fmt.Println(trimmed)
	str1:="Venom"
	str2:="The Last Dance"
    str3:=strings.Join([]string{str1,str2},	"")
    fmt.Println(str3)
}