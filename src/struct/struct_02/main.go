package main

import "fmt"

type person struct{
	Name string
	Age int
}

type doubleZero struct{
	person
	LicenseToKill bool
}

func (p person) greeting(){
	fmt.Printf("this is person's function\n")
}

func (d doubleZero) greeting(){
	fmt.Printf("this is doubleZero's function\n")
}

func main() {
	p1 := person{
		Name : "lmm",
		Age : 19,
	}
	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	
	p1.greeting()
	p2.greeting()
	p2.person.greeting()
}