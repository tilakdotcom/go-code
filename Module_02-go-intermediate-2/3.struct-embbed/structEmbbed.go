package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amout     float32
	status    string
	createdAt time.Time
	Customer
}

// integreting funtion with struct
func (o *Order) changeStatus(status string) {
	o.status = status
}

func main() {
	myOrderWithCustomer := Order{
		id:        "1",
		amout:     50.00,
		status:    "shipped",
		createdAt: time.Now(),
		Customer: Customer{
			name:  "anshu",
			phone: "9876543210",
		},
	}

	fmt.Println(myOrderWithCustomer)

}
