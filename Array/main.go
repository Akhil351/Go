package main
import "fmt"
func main(){
	var name[5] string;
	name[0]="user1"
	name[1]="user2"
	fmt.Println(name)
	var arr=[5]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Println("Length of the array is:",len(arr))
	price:=[5] int{1,2,3,4,4}
	fmt.Println(price)
	var a[5] string;
	a[0]="user1"
	a[1]="user2"
	fmt.Printf("name is %q\n",a)
}