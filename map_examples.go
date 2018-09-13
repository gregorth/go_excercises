package main

import (
	"fmt"
	"bufio"
	"strings"
)

// map an unordered group of elements of one type, called the element type
// indexed by a set of unique keys of another type, called the key type
// The value of uinitilized map is nill
// build on the top of hash table
// map is referenced type

func map_examples() {
	map1 := make(map[string]int)
	map1["jeden"] = 110
	map1["dwa"] = 2
	map1["trzy"] = 3

	fmt.Println("map ", map1)
	// other way of initializing map
	map2 := map[string]int{
		"foo": 1,
		"bar": 2,
		"buu": 3,
		"noo": 4,
	}
	fmt.Println("map", map2)

	x, key_value := map1["jeden"]
	fmt.Println("returned pointer", x, " returned key_value ", key_value)

	k := map1["jeden"]
	fmt.Println("print key", k)
	delete(map1, "jeden")
	fmt.Println("Println map again", map1)

	fmt.Println("map2: ", map2)
}

func map_example_exist() {
	actors := make(map[string]string)
	actors["one"] = "Harrison Ford"
	actors["two"] = "Steven Seagal"
	actors["three"] = "Arnold Szwarceneger"
	actors["four"] = "Mike Jordan"

	if val, exists := actors["five"]; exists {

		fmt.Println("there is a key five")
		fmt.Println("val", val)
	} else {
		fmt.Println("there is no key five")
	}
	// iterating over dict collection
	for key, value := range actors {
		fmt.Println("key: ", key, " value ", value)
	}

}

func scan_words() {
	const input string ="Artificial Intelligence will define the next generation of software solutions. This computer" +
		" science course provides an overview of AI, and explains how it can be used to build smart apps that help" +
		" organizations be more efficient and enrich people’s lives. It uses a mix of engaging lectures and hands-on" +
		" activities to help you take your first steps in the exciting field of AI.Discover how machine learning can" +
		" be used to build predictive models for AI. Learn how software can be used to process, analyze, and extract" +
		" meaning from natural language; and to process images and video to understand the world the way we do. Find" +
		" out how to build intelligent bots that enable conversational communication between humans and AI systems."

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
