package main

import (
	"fmt"
	"time"
)
type Customer struct{
	name string
	number string
}
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer Customer
}

func (o *Order) changeStatus(status string) {
	o.status = status
}

func (o Order) getAmount() float32 {
	return o.amount
}
func newOrder(id string, amount float32, status string) *Order {
	customer:=Customer{
         name:"User",
		 number: "123456789",
	}
	order := Order{
		id:    id,
		amount: amount,
		status: status,
		customer: customer,
	}
	return &order
}

func main() {
	myOrder:=newOrder("1",30.50,"recieved")
	myOrder.customer.name="UserTwo"
	fmt.Println(myOrder)
	// fmt.Println(myOrder.amount)
	// language:= struct{
	// 	name string
	// 	isGood bool
	// }{"Golang",true}
	// fmt.Println(language)
	// order := Order{
	// 	id:     "1",
	// 	amount: 20.00,
	// 	status: "received",
	// }
	// order.changeStatus("confirmed")
	// fmt.Println(order)
	// fmt.Println(order.getAmount())
	// order.createdAt = time.Now()
	// fmt.Println("Order Amount", order.amount)

	// order2 := Order{
	// 	id:        "2",
	// 	amount:    130.00,
	// 	status:    "delieverd",
	// 	createdAt: time.Now(),
	// }
	// order.status = "paid"
	// fmt.Println("Order 1 ", order)
	// fmt.Println("Order 2", order2)

}
