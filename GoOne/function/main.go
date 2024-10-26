package main
import "fmt"
func simpleFunction(){
  fmt.Println("simple Function")
}
func add(a ,b int) (result int) {
	result=a+b
	return
}
func multiply(a,b int) (result int){
	result=a*b
	return
}
func main(){
	fmt.Println("We are Learning function in Golang")
	simpleFunction()
	ans:=add(2,3)
	fmt.Println(ans)
	product:=multiply(2,4)
	fmt.Println(product)
}