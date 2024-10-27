package main

import "fmt"

func getLanguages() (string, string, string,bool) {
	return "Golang", "javaScript", "Java",true
}
func procesIt() func(a int)(int){
      return func(a int) int{
		return a
	  }
}
func main() {
	lang1,lang2,lang3,_:=getLanguages()
	fmt.Println(lang1,lang2,lang3)
	fn:=procesIt()
	fmt.Println(fn(4))
	
}