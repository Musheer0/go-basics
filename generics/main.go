package main

import "fmt"

// func genericType[T any](items []T) {
// 	for i, t := range items {
// 		fmt.Println(i, t)
// 	}
// }
// func main() {
// 	genericType([]any{3,2,1})
// }
func genericType[T comparable](items []T) {
	for i, t := range items {
		fmt.Println(i, t)
	}
}
// func genericType[T bool|string](items []T) {
// 	for i, t := range items {
// 		fmt.Println(i, t)
// 	}
// }

type genericStruct[t any] struct{
	items []t
}

func main() {
	mg := genericStruct[string]{
		items: []string{"m","u"},
	}
	genericType([]int{3,2,1})
}