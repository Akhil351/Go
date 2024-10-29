package main

import "fmt"

//	func emailSender(emailChan <-chan string, done chan<- bool) {
//		defer func() { done <- true }()
//		for email := range emailChan {
//			fmt.Println("sending email", email)
//			time.Sleep(time.Second)
//		}
//	}
func main() {

	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// // closing channel is important
	// close(emailChan)
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recieved data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("recieved data from chan2", chan2Val)
		}

	}

}
