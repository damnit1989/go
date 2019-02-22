package main

import "fmt"

type person struct{
	First string
	Last string
	Age int
}

type doubleZero struct{
	person
	First string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person:person{
			First:"jams",
			Last:"bond",
			Age:20,
		},
		First: "double zero seven",
		LicenseToKill:true,
	}
	p2 := doubleZero{
		person : person{
			First:"Miss",
			Last : "moneyPeny",
			Age : 19,
		},
		First : "if looks could kill",
		LicenseToKill : false,
	}
	p3 := doubleZero{
		person{"Miss","moneyPeny",19},"if looks could kill",false,
	}
	
	fmt.Println(p1.First,p1.person.First)
	fmt.Println(p2.First, p2.person.First)
	fmt.Println(p3.First, p3.person.First)
}
