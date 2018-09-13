package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func(p person) fullname() string {
	return p.firstName +" " + p.lastName
}
func structs_example1() {


	joe := person{"john", "doe", 34}
	fmt.Printf("%T %v \n", joe, joe)
	fmt.Printf("fullname: %s", joe.fullname())

}

type person2 struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person2
	First         string
	LicenseToKill bool
}

func override_fields() {
	p1 := doubleZero{
		person2: person2{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person2: person2{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person2.First)
	fmt.Println(p2.First, p2.person2.First)
}