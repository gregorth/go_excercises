package main

import "fmt"
import "time"

const two_things_in_life string = "tax & death"
const (
	PI                = 3.14
	boiling_water     = 100
	MASTER_OF_DISATER = "Grzegorz"
)

type PizzaSize int

func foo(i int) string {
	fmt.Println("passed arguments", i)
	return "foo"
}

const (
	Small           PizzaSize = 0
	Medium          PizzaSize = 1
	Large           PizzaSize = 2
	ExtraLarge      PizzaSize = 3
	ExtraExtraLarge PizzaSize = 4
)

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

func main() {
	//printing()
	loop_example()
	_ = variadic_example(1, 2, 34, 5, 6)

	fmt.Println("Call variafic function without anyu arguments", variadic_example())
	func_expression_example()

	fmt.Printf("function that return functions its type %T\n", func_return_func())
	fmt.Println("function that return functions", func_return_func()())
	func_with_closure()

	anonymous_self_executing_function()
}
