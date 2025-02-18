package main

import (
	"log"
	"testing"
)

type myStruct struct {
	FirstName string
}

// sign a function to "myStruct" struct this is called receiver
// rather than  accessing the member of the struct directly you can call thi function
// this function can be far more complex than merely returning the value of a given member
// you can perform some business logic in this function
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func TestReceiver(t *testing.T) {
	var myVar myStruct
	myVar.FirstName = "Rezky"

	myVar2 := myStruct{
		FirstName: "Marry",
	}

	log.Println("My Var is set to", myVar.printFirstName())
	log.Println("My Var2 is set to", myVar2.printFirstName())
}
