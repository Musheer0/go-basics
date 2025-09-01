package main

import "fmt"

func add(a, b int, name string) (string,int) {
	fmt.Println(a+b)
	return name,a+b
}

func proces(fn func(a int) int) int{
	return fn(1)
}
func main() {
	name,int:=add(1,2,"mushher")
	fmt.Println(name,int)
}