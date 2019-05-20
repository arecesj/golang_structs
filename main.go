package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// lowercase means private, uppercase means public
// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

// embedding without repetition
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// one way to declare a struct - most popular
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// another way - less popular
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Alex"
	// fmt.Printf("%+v", alex)

	// embed a struct
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Halpert",
	// 	contact: contactInfo{
	// 		email:   "jim@dundermifflin.com",
	// 		zipCode: 94114,
	// 	},
	// }
	// fmt.Printf("%+v", jim)
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim@dundermifflin.com",
			zipCode: 94114,
		},
	}
	// verbose way of using pointers to update a struct
	// jimPointer := &jim
	// jimPointer.updateName("Dwight")
	// jim.print()

	// Golang is OK with interpreting this as a pointer when you invoke the fn
	// bc of the *person in the updateName fn
	jim.updateName("Michael")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
