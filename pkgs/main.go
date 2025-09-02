package main

import (
	auth "github.com/Musheer0/user/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("musheer","12345")
	auth.GetSession()
	user := auth.User{
		Name: "musheer",
		
	}
	color.Red(user.Name)
}