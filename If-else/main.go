package main
import "fmt"
func main(){
	x:=10
	if x>5{
		fmt.Println("X is greater than 5")
	} else{
		fmt.Println("X is less than an equal to 5")
	}
	z:=10
	if z>10{
		fmt.Println("z is greater than 10")
	} else if(z>5){
		fmt.Println("z is greater 5 and smaller than 10")
	}else{
		fmt.Println("z is smaller than 5")
	}
}