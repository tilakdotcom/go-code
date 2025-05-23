package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amout     float32
	status    string
	createdAt time.Time
}

// integreting funtion with struct
func (o *Order) changeStatus(status string) {
	o.status = status
}

func (o Order) getAmmout() float32 {
	return o.amout
}

// A convension to use struct for easiness
func newOrder(id string, amount float32, status string) *Order {
	order := Order{
		id:        id,
		amout:     amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &order
}

func main() {
	//if u won't set any field then it uses zeroed values by default
	myOrder1 := Order{
		id:        "1",
		amout:     50.00,
		status:    "shipped",
		createdAt: time.Now(),
	}

	myOrder1.status = "deliverd"

	myOrder2 := Order{
		id:        "2",
		amout:     40.00,
		status:    "paid",
		createdAt: time.Now(),
	}
	fmt.Println("first order ", myOrder1)
	fmt.Println("second order ", myOrder2)

	//struct with funtions
	myOrder3 := Order{
		id:        "3",
		amout:     40.00,
		status:    "shiiped",
		createdAt: time.Now(),
	}

	myOrder3.changeStatus("confirm")

	fmt.Println("order 3", myOrder3)

	fmt.Println("order 3 amount", myOrder3.getAmmout())

	// convesional way to create order
	myOrder4 := newOrder("4", 20.00, "shipped")

	fmt.Println("forth order ", myOrder4)

	// inline way to create struct
	language := struct {
		name   string
		isGood bool
	}{
		name:   "Go lang",
		isGood: true,
	}

	fmt.Println("inline struct ", language)

}
