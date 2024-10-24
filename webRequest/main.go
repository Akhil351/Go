package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func main() {
	fmt.Println("Learning web Request")
	// Making web request
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if(err!=nil){
		 fmt.Println("Getting error in get Response",err)
		 return
	}
	defer res.Body.Close()
	data,_:=ioutil.ReadAll(res.Body)
	if(err!=nil){
		fmt.Println("Error")
		return
	}
	fmt.Println(string(data))

}


// When you make an HTTP request with http.Get(), 
// Go establishes a network connection and keeps it open to read the response body.
//  If you don't explicitly close the response body, the connection stays open 
//  and resources are not released, which can lead to memory leaks or resource exhaustion if you're making many requests.