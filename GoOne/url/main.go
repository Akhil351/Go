package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url")
	myUrl:="https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("url type is %T\n",myUrl)
	parsedUrl,err:=url.Parse(myUrl)
	if(err!=nil){
		fmt.Println("Can't parse url",err)
	}
	fmt.Printf("url type is %T\n",parsedUrl)
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)

}
