package main
import "fmt"
const age=23
func main(){
	const name string="golang"
	const (
		port=5000
		host="localhost"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port,host)
}