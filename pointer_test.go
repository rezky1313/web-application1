package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var myString string
	myString = "Green"

	fmt.Println("myString is set to", myString)
	ChangeUsingPointer(&myString)
}

func ChangeUsingPointer(s *string)  {
	newValue := "red"
	s = &newValue
}