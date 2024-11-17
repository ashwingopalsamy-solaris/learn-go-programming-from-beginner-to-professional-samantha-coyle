package main

import "fmt"

type Patient struct {
	firstName  string
	secondName string
	age        int
	pAllergy   bool
}

func main() {
	p1 := Patient{
		firstName:  "Ashwin",
		secondName: "Gopalsamy",
		age:        24,
		pAllergy:   false,
	}
	fmt.Printf("%s\n%s\n%d\n%t\n", p1.firstName, p1.secondName, p1.age, p1.pAllergy)
}
