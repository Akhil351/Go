package main
import "fmt"
func main(){
	day:=5
	switch day{
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesaday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}

	month:="December"
	switch month{
	case "January","February","March":
		fmt.Println("Winter")
	case "April","May","June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}
	
}