package main

import (
	"fmt"
	"maps"
)

func main() {
	//basic
	// m := make(map[string]string)
	// m["key"] = "value"
	// m["de"]="le"
	// delete(m,"de")
	// fmt.Print(m["key"])
	// k,ok := m["key"]
	// fmt.Print(k,ok)
	// clear(m)

	m1 :=map[string]int{"price":565,"stock":2}
	m2:=map[string]int{"price":545,"stock":44}
	fmt.Print(maps.Equal(m1,m2))

	nums:=[]int{1,2}

	for _,num:=range nums{
		fmt.Print(num)
	}
	for i,c:=range "goland"{
		fmt.Println(i,c)
	}
}