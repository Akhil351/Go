package main
import "fmt"
func add(a,b int) (result int) {
	result=a+b
	return
}
func main(){
	fmt.Println("Start of the Program")
	data:=add(23,23)
	defer fmt.Println("data is ",data)
    defer fmt.Println("Middle of the Program")
	fmt.Println("End of the Program")
}