// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string{
	return "This is the sample function"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string{
	return "Hello, My name is " + name + "!"
}

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) string{
	return "Hello, My name is " + name + "! and my age is " + strconv.Itoa(age)+"."
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(a,b int)(int, error){
	if(b==0){
		return 0, errors.New("Divided by Zero error")
	}
	return a/b, nil
}

func main() {
	// invoke greet function
	// fmt.Println(greet())

	// invoke greetWithName function
	// fmt.Println(greetWithName("Rakesh"))
// fmt.Println(greetWithNameAndAge("Rakesh", 33))
	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
