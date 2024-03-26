// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum
func TestSum(t *testing.T) {
	actual := sum(1, 2, 3)
	expected := 6
	if actual != expected {
		t.Errorf("Actual value %v and Expected value %v", actual, expected)
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M){
	fmt.Println("Before starting the Test")
	m.Run()
	fmt.Println("After the Test")
}