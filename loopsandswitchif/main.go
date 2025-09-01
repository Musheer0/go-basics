package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		i++
		if i==3{
			break
		}
		if i==2{
			continue
		}
		fmt.Print(i)
		fmt.Printf("\n")
	}
	var names  = [2]string{"hello","world"} 
	fmt.Print(len(names))
	// for i:=range(55){
	// 	println(i)
	// }
	typecheck :=func (i interface{})  {
		switch i.(type){
		case string:
			fmt.Printf("this is string")
		case int:
			fmt.Printf("This is integer")
		case bool:
			fmt.Printf("This is booleamn")
		case float32:
			fmt.Printf("This is float32")
		}
	}
	typecheck(true)
}