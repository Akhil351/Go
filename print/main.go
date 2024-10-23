package main
import ("fmt")
func main(){
	age:=25
	name:="Alice"
	height:=5.122321
	fmt.Println("age:",age,"name:",name,"height:",height)
	fmt.Print("age:",age,"name:",name,"height:",height)
	fmt.Println()
	fmt.Printf("age is %d\n",age)
	fmt.Printf("height is %.3f\n",height)
	fmt.Printf("Type of name is %T\n",name)
	fmt.Printf("Type of name is %T\n",age)
	fmt.Printf("Type of name is %T\n",height)
	fmt.Printf("name is %s\n",name)
	fmt.Printf("name is %s, Age is %d , height is %.2f\n",name,age,height)
}