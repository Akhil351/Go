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

	temperatue:=25
	switch {
	case temperatue<0:
		fmt.Println("freezing")
	case temperatue>=0 && temperatue<10:
		fmt.Println("Cold")
	case temperatue>=10 && temperatue<20:
		fmt.Println("Cool")
	case temperatue>=20 && temperatue<30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
	
}