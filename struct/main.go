package main

import (
	"fmt"
	"time"
)

type order struct {
	id         string
	amount     float32
	status     string
	created_at time.Time
	customer_id string
}


//add a function to struct
func (o *order) changeStatus(status string){
	o.status =status
}
//where struct dosent have any contructor but we can ductape it using this method
//here  func name can be anything but  adding prefix new makes it easier to filter out with other helper function
func newOrder(o order)  *order{
	myorder := order{
		id:o.id,
		status: o.status,
		created_at: o.created_at,
		//lets add 15% tax here for example
		amount: o.amount+(o.amount*0.15),
		//example we add from session token
		//uuid token with user_ prefix
		customer_id: "user_bc328fca-4ec2-48d7-b313-206dbe7b9d99",
		
	}
	
	return  &myorder
}

//nested struct by the name itself you should know the usage
//here is an example
type customer struct{
	name string
	age int
	is_plus_member bool
	orders  []order
	order_count int
}

//lets create a intilizer for it to calc in total number of orders
func newCustomer(c customer) *customer{
	mycustomer := c
	mycustomer.order_count = len(c.orders)
	return  &mycustomer
} 

func main() {	
	// or:=order{
	// 	id:"23DQ",
	// 	amount: 23.4,
	// 	status: "otw",
	// 	created_at: time.Now(),
	// }
	
	// or.changeStatus("out for delivery")
	// fmt.Println(or)

	//crating a inline struct like these can be used when we have to use that struct only one
	//you can create using struct keyword and inside {} pass the types and open a second {} where you need to pass the arg in same order of types
	// inlineStruct :=struct{
	// 	name string
	// 	age int
	// }{"musheer",17 }
	// fmt.Println(inlineStruct)

	initializedOrder :=*newOrder(
		order{
		id:"23DQ",
		amount: 23.4,
		status: "otw",
		created_at: time.Now(),
	})
	initializedOrder.changeStatus("out for delivery")
	//works like normal struct
// fmt.Println(initializedOrder)
	mycustomer := *newCustomer(
		customer{
			name: "musheer",
			is_plus_member: true,
			orders: []order{initializedOrder},
		})
	fmt.Println(mycustomer)
}