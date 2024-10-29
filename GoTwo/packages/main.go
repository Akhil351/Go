package main

import (
	"akhil/auth"
	"akhil/user"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("akhil", "secret")
	session := auth.GetSession()
	fmt.Println(session)
	user := user.User{
		Email: "akhil.vathaluru@gmail.com",
		Name:  "akhil",
	}
	fmt.Println(user)
	color.Green(user.Email)
	color.Blue(user.Name)
}
