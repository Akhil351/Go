package main
import "fmt"
func divide(a,b float64) (float64,error){
	if(b==0){
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a/b,nil
}
func main(){
	ans,err:= divide(10,6) // ans,_:=divide(10,0)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("Divison of two numbers is ",ans)
	}
	
}