package main

import (
	"log"
	"testing"
)

type UserMap struct {
	FirstName string
	LastName  string
}

func TestMapAndSlices(t *testing.T) {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	log.Println(myMap["dog"])

	myMap["other-dog"] = "Cassie"
	log.Println(myMap["other-dog"])

	//overwrite the value of the key "dog"
	myMap["dog"] = "Miko"
	log.Println(myMap["dog"])

	//map using struct
	myMap2 := make(map[string]UserMap)

	me := UserMap{
		FirstName: "Rezky",
		LastName:  "Wahyudi",
	}

	myMap2["me"] = me

	log.Println(myMap2["me"].FirstName)
	log.Println(myMap2["me"].LastName)
}

func TestSlices(t *testing.T) {
	var myInt []int

	myInt = append(myInt, 1)
	myInt = append(myInt, 2)

	log.Println(myInt)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers[2:5])
}
