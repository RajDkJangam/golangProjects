package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//zalak := person{firstName: "Zalak", lastName: "Shah"}
	zalak := person{
		firstName: "Zalak",
		lastName:  "Shah",
		contact: contactInfo{
			email:   "zalak13@gmail.com",
			zipCode: 02115,
		},
	}
	//var zalak person
	zalak.firstName = "Zalak"
	zalak.lastName = "Shah"
	fmt.Println(zalak)
	fmt.Printf("%+v\n", zalak)
}
