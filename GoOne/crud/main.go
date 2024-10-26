package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Compelted bool   `json:"compelted"`
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting: ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting the response: ", res.Status)
		return
	}
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding:", err)
		return
	}
	fmt.Println("Todo", todo)
	fmt.Println(todo.Compelted)
	fmt.Println(todo.UserId)
}
func postRequest() {
	todo := Todo{
		UserId:    351,
		Title:     "PostMethod in Go lang",
		Compelted: true,
	}
	// convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	// convert json data  to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending the request :", err)
		return
	}
	defer res.Body.Close()
	// data,_:=ioutil.ReadAll(res.Body)
	// fmt.Println(string(data))
	fmt.Println("Response Status", res.Status)

}
func putRequest(){
	todo:=Todo{
		UserId: 23789,
		Title: "Updating the request",
		Compelted: false,
	}

	// convert to the json
	json,err:=json.Marshal(todo)
	if(err!=nil){
		fmt.Println("Error marshalling",err)
		return
	}
	// convert json data to string
	jsonString:=string(json)

	// convert string to reader
	jsonReader:=strings.NewReader(jsonString)

	const myUrl="https://jsonplaceholder.typicode.com/todos/1"
	// create the put request

	req,err:=http.NewRequest(http.MethodPut,myUrl,jsonReader)
	if(err!=nil){
		fmt.Println("Error Creating Put Request: ",err)
		return
	}
	req.Header.Set("Content-type","application/json")

	client:=http.Client{}
	res,err:=client.Do(req)
	if(err!=nil){
		fmt.Println("Error Sending the request: ",err)
		return
	}
	defer res.Body.Close()
	data,_:=ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}
func deleteRequest(){
	const myUrl="https://jsonplaceholder.typicode.com/todos/1"
	// create delete request

	req,err:=http.NewRequest(http.MethodDelete,myUrl,nil)
	if(err!=nil){
		fmt.Println("Error Creating Put Request: ",err)
		return
	}
	req.Header.Set("Content-type","application/json")

	client:=http.Client{}
	res,err:=client.Do(req)
	if(err!=nil){
		fmt.Println("Error Sending the request: ",err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res.Status)

}
func main() {
	fmt.Println("Learning Crud in Go Lang")
	//getRequest()
	//postRequest()
	//putRequest()
	deleteRequest()

}
