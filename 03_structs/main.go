package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

// Create embedded structs
type contactInfo struct {
	email   string
	zipCode int
}

type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Two ways of instantiating structs
	p := person{"Matias", "Borghi"}
	fmt.Println(p)
	p = person{firstName: "Someone", lastName: "Else"}
	fmt.Println(p.firstName, p.lastName)

	// zero value assigns default values
	// string -> ""
	// int -> 0
	// float -> 0
	// bool -> false
	var zeroValue person
	fmt.Println(zeroValue)         // prints { }
	fmt.Printf("%+v\n", zeroValue) // prints {firstName: lastName:}

	zeroValue.firstName = "Matias"
	zeroValue.lastName = "Borghi"
	fmt.Printf("%+v\n", zeroValue) // prints {firstName:Matias lastName:Borghi}

	p2 := person2{
		firstName: "Matias",
		lastName:  "Borghi",
		contact: contactInfo{
			email:   "borghi.matias at email.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v\n", p2) // prints {firstName:Matias lastName:Borghi contact:{email:borghi.matias at email.com zipCode:12345}}

	// This is not updated because p2 is passed as value. So the copy is updated.
	// Instead, we have to pass the p2 reference.
	// p2.updateName("Jim")

	// Using this instead updates the name correctly
	// &var gives the memory address
	// p2Pointer := &p2
	// p2Pointer.updateFirstName("Jim")
	// We can also directly call the receiver function to the value avoid code
	p2.updateFirstName("Jim")
	p2.print()
}

// Define a receiver function for person2
func (p person2) print() {
	fmt.Printf("%+v\n", p)
}

// func (p person2) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerP *person2) updateFirstName(newFirstName string) {
	// *var is the value of the memory address that is pointing at
	(*pointerP).firstName = newFirstName
}
