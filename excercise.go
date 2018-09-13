package main

import (
	"fmt"
	"sort"
)

type people []string

// Sort implementation for []string
func(p people) Len() int{
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func sort1() {
	studyGroup := people{"John", "Zeno", "Al", "Jenny"}
	// sort people by name
	// then sort it reverse
	fmt.Println("Before sort: ", studyGroup)
	// jesli typ implementuje interfesj to sie dzieje automatycznie przez receiver
	// trzeba napisac receiver dla sorta
	// we need to imlement three function Len, Less, Swap
	sort.Sort(studyGroup)
	fmt.Println("After sort: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("Sorted reverse: ", studyGroup)


}
// Sort implementation for []string

//func (s []string) Len() int {
//
//}

func sort2() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Before sort:", s)
	sort.Strings(s)
	fmt.Println("After sort:", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("After sort reverse", s)
}

func sort3() {
	// sort in normal and reverse order
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println("Integers sorted", s)
	fmt.Println("list is sorted?:",sort.IntsAreSorted(s))
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//sort.Reverse(sort.IntSlice(s))
	fmt.Println("Integers sorted in a reverss way", s)
}

func main() {
	//fmt.Println("1 excercise")
	//sort1()
	//fmt.Println("2 excercise")
	//sort2()
	//fmt.Println("3 excercise")
	sort3()
}