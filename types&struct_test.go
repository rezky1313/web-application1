package main

import (
	"log"
	"testing"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	age int
	BirthDate time.Time
}

func TestTypeAndStruct(t *testing.T) {
	user := User{
		FirstName: "Rezky",
		LastName: "Wahyudi",
	}

	log.Println(user.FirstName, "BirthDate", user.BirthDate)
}

