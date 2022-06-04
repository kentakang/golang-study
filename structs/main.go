package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{}

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
}
