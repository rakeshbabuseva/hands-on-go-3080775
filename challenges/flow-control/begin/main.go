// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func ()  {
		if err:= recover(); err!=nil{
			fmt.Println("recovered from panic ", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	filename := os.Args[1]
fmt.Println(filename)
	data, err := os.ReadFile(filename)

	if err!=nil{
		panic("Unable to read the file ")
	}


	// convert the bytes to a string
	data2 := string(data)
	// initialize a map to store the counts
	counter := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	dataInArray := strings.Fields(data2)

	// capture the length of the words slice
	totalNoOfWords := len(dataInArray)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
counter["words"] = totalNoOfWords
	for _, words := range dataInArray{
		for _, char := range words{
			if unicode.IsNumber(char){
				counter["numbers"] = counter["numbers"] + 1
			} else if unicode.IsLetter(char) {
				counter["letters"] = counter["letters"] + 1
			} else {
				counter["symbols"] = counter["symbols"] + 1
			}
		}
	}
	spew.Dump(counter)

	// dump the map to the console using the spew package
	//
}
