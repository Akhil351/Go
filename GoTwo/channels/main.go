package main

import "fmt"

// send the data
//	func processNum(numChan chan int) {
//		fmt.Println("processing number", <-numChan)
//	}

// recieve the data
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult

// }

// go routine sychronizer
func task(done chan bool) {
	defer func(){ done <- true}()
	fmt.Println("processing...")
}
func main() {
	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 6

	//	time.Sleep(time.Second)

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result
	// fmt.Println(res)



	done:=make(chan bool)
	go task(done)
	<- done

}
