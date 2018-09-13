package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

type Person2 struct {
	First string
	Last string
	Age int `json:"wisdom score"`
}
func json1 () {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	err := json.Unmarshal(bs, &p1)
	fmt.Println(err)
	fmt.Println("---")
	fmt.Println("first: ", p1.First)
	fmt.Println("last: ", p1.Last)
	fmt.Println("age: ", p1.Age)

	fmt.Printf("%T %v \n", p1, p1)
}

func json2() {
	var p2 Person2

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	err := json.Unmarshal(bs, &p2)
	fmt.Println(err)
	fmt.Println("---")
	fmt.Println("first: ", p2.First)
	fmt.Println("last: ", p2.Last)
	fmt.Println("age: ", p2.Age)
}
