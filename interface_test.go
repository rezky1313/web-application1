package main

import (
	"fmt"
	"testing"
)

type Animal interface {
	//to use an interface you must have a function that match this spec
	//for example this function below named Says and returning a string, name of th function can be random
	Says() string
	NumberOfLegs() string
	AnimalName() string
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

var DogVar Dog

func TestInterface(t *testing.T) {

	PrintInfo(&DogVar)

	fmt.Println(DogVar.Name)
}

// method for type Interface
func PrintInfo(a Animal) {
	fmt.Println("this Animal Says", a.Says(), "and Has", a.NumberOfLegs(), "Legs", "and named", a.AnimalName())
}

func (d *Dog) Says() string {
	return "Woorf"
}

func (d *Dog) NumberOfLegs() string {
	return "4"
}

func (d *Dog) AnimalName() string {
	d.Name = "Samson"
	return d.Name
}
