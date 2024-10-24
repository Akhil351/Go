package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error whil creating file:", err)
	// }
	// defer file.Close()
	// content:="Venom the last dance only theaters now"
	// _,errors:=io.WriteString(file,content+"\n")
	// if(errors!=nil){
	// 	fmt.Println("Error while writing the file:",err)
	// 	return 
	// }
	// fmt.Println("SuccessFully Created the file")
	file,err:=os.Open("example.txt")
	if(err!=nil){
		fmt.Println("Error while Opening the file:",err)
	}
	defer file.Close()
	

}
