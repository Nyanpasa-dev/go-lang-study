package main

import (
	"fmt"
)

func main() {
	receiveCredentials("Max", 15)
	receiveCredentials("Baha Box", 32)
	receiveCredentials("Sanya Kick Box", 15)
}

func print(message string) {
	fmt.Println(message)
}

func receiveCredentials(name string, age int) {
	if checkAge(age) {
		print(fmt.Sprintf("Name: %s, Age, %d", name, age))
	} else {
		print("You are not allowed to enter")
	}
}

func checkAge(age int) bool {
	if age < 18 {
		return false
	}

	return true
}
