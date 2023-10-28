package main

import "fmt"

func main() {

	n := NewLoginDialog()

	n.username.setContent("username")
	n.password.setContent("password")
	n.agreeToTerms.setEnable(true)

	fmt.Printf("signupBotton.isEnabled: %v\n", n.signupBotton.isEnabled)

	n.username.setContent("username")
	n.password.setContent("password")
	n.agreeToTerms.setEnable(false)

	fmt.Printf("signupBotton.isEnabled: %v\n", n.signupBotton.isEnabled)
}
