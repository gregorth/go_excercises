package main

import "fmt"
import "reflect"

// multidimensial slice just like excel spreadsheet
func multidimensial_slice() {
	records := make([][]string, 0)
	// first row
	row := make([]string, 4)
	row[0] = "John"
	row[1] = "Doe"
	row[2] = "36 years old"
	row[3] = "8086"
	// append row to records
	records = append(records, row)
	row2 := make([]string, 4)
	row2[0] = "Joanna"
	row2[1] = "Krupa"
	row2[2] = "39 years old"
	row2[3] = "8087"
	// append
	records = append(records, row2)

	fmt.Println(records)
}

func dynamic_multidimensial(rows_number int, cols_number int) {
	matrix := make([][]int, 0)

	for i := 0; i < rows_number; i++ {
		// make([]T, len, cap) => []T
		// way of solve 1
		// column := make([]int, cols_number)
		column := make([]int, 0, 0)
		for j := 0; j < cols_number; j++ {
			// column[j] = j * i
			column = append(column, i*j)
		}
		matrix = append(matrix, column)
	}
	fmt.Println("created matrix with rows: ", rows_number, " and columns ", cols_number)
	fmt.Println(matrix)
}

func slice_creation() {
	// shorthand
	var student []string //got underlying array, is not null
	var students [][]string
	// var
	student_short := []string{} // var setup zero value for type  =>
	// no underlying slice is created so it
	// will be eqaul nil
	students_short := [][]string{}
	// by make
	class_of_students := make([]string, 0, 0)

	fmt.Println("Slice creation")
	fmt.Println("student ", student)
	fmt.Println("students", students)
	fmt.Println("student == nil", student == nil)
	fmt.Println("student_short", student_short)
	fmt.Println("students_short", students_short)
	fmt.Println("students_short == nil", students_short == nil)
	fmt.Println("Class of students", class_of_students)
	fmt.Println("Class of students == nil", class_of_students == nil)
	// delete elements from slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	digits := append([]int(nil), numbers...)
	fmt.Println("Numbers ", numbers)
	fmt.Println("Digits is the copy of numbers", digits)
	// comparing two slices
	fmt.Println("numbers == digits", reflect.DeepEqual(numbers, digits))
	// remove 4th and 6th element from numbers
	// result should be [1, 2, 3, 4, 6, 8]
	numbers = append(numbers[0:4], numbers[5], numbers[7])
	fmt.Println("numbers with removed elements 4th and 6th", numbers)
	fmt.Println("numbers == digits", reflect.DeepEqual(numbers, digits))
	numbers2 := []int{}

	copy(numbers2, digits)
	fmt.Println("sliced copied using copy function", numbers2)
	fmt.Println("numbers2 == digits", reflect.DeepEqual(numbers2, digits))
	//it seems that copy doesn work when target slice doesn't have length equal source slice
	numbers3 := make([]int, 3, 3)
	copy_count := copy(numbers3, digits)
	fmt.Println("content of numbers3 slice", numbers3)
	fmt.Println("amount of copied elements", copy_count)
	// slicing array returns slice

}
