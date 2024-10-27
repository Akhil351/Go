package main

import "fmt"

type Paymenter interface {
	pay(amount float32)
	refund(amount float32,account string)
}
type Payment struct {
	gateway Paymenter
}

func (p Payment) makePayment(amount float32) {
	p.gateway.pay(amount)
	p.gateway.refund(amount,"123")
}

type razoPay struct{}

func (r razoPay) pay(amount float32) {
	fmt.Println("making payment using razoPay", amount)
}
func (r razoPay) refund(amount float32,account string) {
	fmt.Println("making refund using razoPay", amount,account)
}

type Stripe struct{}

func (s Stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}


func main() {
	//stripe:=Stripe{}
	razoPay := razoPay{}
	payement := Payment{
		gateway: razoPay,
	}
	payement.makePayment(100)
}
