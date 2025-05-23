package main

import "fmt"

type payment struct {
	gateway paymenter
}
type razorpay struct{}
type stripe struct{}
type fakePayment struct{}

type paymenter interface {
	pay(amout float32)
}

func (r razorpay) pay(amout float32) {
	//login of razorpay payment
	fmt.Println("making payment using razorpay ", amout)
}

func (f fakePayment) pay(amout float32) {
	//login of razorpay payment
	fmt.Println("making payment using FakeGW ", amout)
}

func (s stripe) pay(amout float32) {
	//login of razorpay payment
	fmt.Println("making payment using stripe ", amout)
}

func (p payment) makePayment(amout float32) {
	p.gateway.pay(amout)
}

func main() {
	newPayment := payment{
		// gateway: razorpay{},
		gateway: stripe{},
		// gateway: razorpay{},
	}
	newPayment.makePayment(50.00)
}
