package main
import "fmt"
func main(){
	for i:=0;i<5;i++{
		fmt.Println(i);
	}
	counter:=0
	for{
		fmt.Println("Infinte Loop")
		counter++
		if(counter==3){
			break
		}
	}
	numbers:=[]int{1,2,3,4,5}
	for index,value:= range numbers{
		fmt.Printf("index %d ,value %d\n",index,value)
	}
	data:="Hello World"
	for index,char:=range data{
		fmt.Printf("index is %d,character is %c\n",index,char)
	}
}