package main

import "fmt"

type status int

const (
	online status = iota
	offline
	busy
	disabled
)

// type status string

// const (
// 	online status = "online"
// 	offline status        ="offilne"
// 	busy    status       = "busy"
// 	disabled  status     = "disabled"
// )
func printStatus(s status) {
	fmt.Println("user is",s)
}
func main() {
	printStatus(online)
}