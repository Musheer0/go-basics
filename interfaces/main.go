package main

import "fmt"

type ipayment interface {
	pay(amt float32)
	refunc(amt float32,accnt string)
}


type payment struct {
	gateway ipayment
}

func (p payment) paymoney (amt float32){
	p.gateway.pay(amt)
}

type paypal struct{}

func (p paypal) pay(amt float32) {
	fmt.Println("method:paypal", amt)
}
func (p paypal) refunc(amt float32,acct string){
	fmt.Println("refund\nmethod paypal amount:",amt," account id ",acct)
}
type razorpay struct{}

func (p razorpay) pay(amt float32) {
	fmt.Println("method:razorpay", amt)
}

func main() {

	newpayment:=payment{
		gateway: paypal{},
	}
	newpayment.paymoney(45.66)
	newpayment.gateway.refunc(33.22,"user_e3jd3s2")
}