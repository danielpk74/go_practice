package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	claudia := person{
		firstName: "Claudia",
		lastName:  "Londoño",
		contactInfo: contactInfo{
			email:   "clau@email.com",
			zipCode: 94000,
		},
	}

	//fmt.Printf("%+v", claudia)
	claudia.print()
	claudia.updateName("Claudia María")
	claudia.print()

	// daniel := person{"Daniel", "Montoya"}
	// var claudia person
	// andrea := person{firstName: "Andrea", lastName: "Atencio"}

	// fmt.Println(daniel.firstName, claudia, andrea.firstName)
}
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateNameWihoutPointers(newFirstName string) {
	p.firstName = newFirstName
}
