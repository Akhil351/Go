package main
import ("fmt")
func main(){
	numbers:=[]int{1,2,3,4,5}
	numbers=append(numbers, 6,7,8,9,10)
	fmt.Println(numbers)
	fmt.Printf("data Type %T\n",numbers)
	fmt.Println("length of the array:",len(numbers))
	name:=[]string{}
	fmt.Println(name)
	fmt.Println("length",len(numbers))
	fmt.Println("capcacity",cap(numbers))

	a:=make([] int,3,5)
	a=append(a, 4)
	a=append(a, 5,6,7,8,9)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}