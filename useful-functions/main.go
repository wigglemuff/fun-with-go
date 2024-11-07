package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

var p = fmt.Println

func main() {

	// --------------------------------------------------
	// Sorting a slice of strings.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// --------------------------------------------------
	// Sorting a slice of integers.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// --------------------------------------------------
	// Checking if a slice is sorted.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)

	// --------------------------------------------------
	// Sorting with a custom function.
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// --------------------------------------------------
	// Sorting a struct with a custom function.
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)

	// --------------------------------------------------
	// Common string functions
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))

	// Contains:   true
	// Count:      2
	// HasPrefix:  true
	// HasSuffix:  true
	// Index:      1
	// Join:       a-b
	// Repeat:     aaaaa
	// Replace:    f00
	// Replace:    f0o
	// Split:      [a b c d e]
	// ToLower:    test
	// ToUpper:    TEST

	// --------------------------------------------------
}
