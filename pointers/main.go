package main

import "fmt"
//* for pointer pointing to that address
func printNum(num *int) {
	*num = 5
	fmt.Println("change",*num)
}

func main(){
	num:=1
	fmt.Println("before",num)
	//& to get the adress
	printNum(&num)
	//mem addresss
	//prefix & to get mem address
	fmt.Println("after",num)
}