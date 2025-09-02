package auth

import "fmt"

func LoginWithCredentials(name string, password string) {
	fmt.Println(name,password)
}
type User struct{
	Name string
	email string
}