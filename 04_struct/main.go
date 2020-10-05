package main

import "fmt"

// structure is like Python dictionary

// define struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.print()
}

// pass by value. refer pass_by_Value.png
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// pass by reference
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// receiver name p should be consistent with previous receiver name pointerToPerson for person
func (p person) print() {
	//fmt.Println(p)
	fmt.Printf("%+v", p)
}
