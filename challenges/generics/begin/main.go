// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring
func printAny[T int | string | bool](i T) {
	fmt.Println(i)
}

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
type numeric interface {
	constraints.Integer | constraints.Float
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T

	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	// printAny[string]("Hello")
	// printAny[int](42)
	// printAny[bool](true)
	// call sum function
	fmt.Println("result", sum(1.4, 2.3, 3.5))

	// call generics sumAny function
}
