package main

import (
	"fmt"
	"time"
)

const two_things_in_life string = "tax & death"
const (
	PI                = 3.14
	boiling_water     = 100
	MASTER_OF_DISATER = "Grzegorz"
)
const (
	Small           PizzaSize = 0
	Medium          PizzaSize = 1
	Large           PizzaSize = 2
	ExtraLarge      PizzaSize = 3
	ExtraExtraLarge PizzaSize = 4
)

type PizzaSize int

func printing() {

	now := time.Now()
	year, week_of_the_year := now.ISOWeek()
	fmt.Printf("hello, world\n")
	fmt.Printf("Today it is %v which is %v day and %v week of year %v\n",
		now.Format(time.UnixDate),
		now.YearDay(),
		week_of_the_year,
		year,
	)
	x := 42
	var name string = "G ś"
	fmt.Printf("My name abbrevations %v", name)
	fmt.Println("calling a function...", foo(4))
	const GGG = 42
	fmt.Print("SSSSSSSSSSSSSSSSS")
	fmt.Print("X: ", x)
	fmt.Println("X adresss ", &x)
	fmt.Println("ttil", two_things_in_life)
	fmt.Println(GGG)
	// fmt.Println(year)
	fmt.Println("Q - ", GGG)
	fmt.Println(PI)
	fmt.Println(MASTER_OF_DISATER)
	fmt.Println(boiling_water)
	// fmt.Println(week_of_the_year)
	// fmt.Println("wow should print something")
	// fmt.Println(now.Weekday())
	fmt.Println("Pizza size extralarge", ExtraExtraLarge)
}

func loop_example() {
	println("loop example")
	sum := 0
	for i := 0; i < 12; i++ {
		sum += i
	}
	println("Final sum:", sum)
}

func variadic_example(numbers ...int) float64 {
	var sum float64
	for _, v := range numbers {
		sum += float64(v)
	}
	fmt.Println("summed", sum)
	return sum
}

func func_expression_example() {

	count := func(numbers ...int) int {
		return len(numbers)
	}
	fmt.Println(" count of list (1,2,3)", count(1, 2, 3))
	fmt.Println(" count of list (1,2,3,4,5 ,6 ,7)", count(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(" count of list ()", count())
	fmt.Printf("func expression type %T\n", count)
}

func func_return_func() func() string {
	return func() string {
		return "Hello WORLD"
	}
}

func func_with_closure() {
	x := 0
	inc := func() int {
		x++
		return x
	}
	inc()
	inc()
	fmt.Printf("%v", x)
}

func func_with_closure2() {
	y := 0
	inc := func() int {
		y++
		return y
	}
	inc()
	inc()
	fmt.Printf("%v", y)
}

func anonymous_self_executing_function() {
	func() {
		fmt.Println("Self executing... Done")
	}()
}

func foo(i int) string {
	fmt.Println("passed arguments", i)
	return "foo"
}

//function verification
//   expected
//   is_even(1) returns (0, false)
//   is_even(2) returns (1, true)
func is_even(a int) (int, bool) {
	var is_even bool
	if a%2 == 0 {
		is_even = true
	} else {
		is_even = false
	}
	return a / 2, is_even

}

// the same as above but with func expression
func is__even2(b int) (int, bool) {
	is_even2 := func(a int) (int, bool) {
		is_even := false
		if a%2 == 0 {
			is_even = true
		} else {
			is_even = false
		}
		return a / 2, is_even
	}
	return is_even2(b)
}

/*
IMPORTANT
Avoid using named returns.
Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/

// find_max([1, 22, 12, 9 ,34]) => 34
// proper way of calling it:
// 	nums := []int{1, 22, 89, 12, 34, 7, 91}
//  find_max(nums...)
func find_max(numbers ...int) int {
	max := 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}
