package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

/*
Json stands for javascript notation and its just way of getting information from one place to another
in a structured format. And go makes it really easy to work with
*/

// assign the json file to our data structure
// we can use this type to unmarshal json we got from an external API into a struct
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func TestJson(t *testing.T) {
	//json package we pretend got from an external API, it might be thousand of data
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	//unmarshall the json package
	var unmarshalled []Person

	//unmarshal function from golang took 2 parameters wich slice of byte wich is strings from json package wich is myjson variable
	//and an interface that youre going to put the contents if the slice of byte into and thats will be variable
	//that we created before called unmarshalled
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Erorr Unmarshalling Json")
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	//so we can read Json into a struct

	//now how do we write Json from a struct?
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "Nagari"
	m1.HairColor = "Red"
	m1.HasDog = false

	//append the variable into the slice "mySlice"
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "John"
	m2.LastName = "Kennedy"
	m2.HairColor = "Brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	log.Println(mySlice)

	//so now  the slice consist two entries
	//now assigne the slice of dta into a json

	newJson, err := json.MarshalIndent(mySlice, "This is my first Json YO", "     ")
	if err != nil {
		log.Println("Error Marshalling the data", err)
	}

	//print json we created, but first convert it into a string since it convert into a byte
	fmt.Println(string(newJson))

}
