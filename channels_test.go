package main

import (
	"fmt"
	"testing"

	"github.com/rezky1313/web-application1/helpers"
)

// channels mean sending information from one part of your program to another part of your program
func TestChannel(t *testing.T) {
	intChantek := make(chan int)
	defer close(intChantek)

	for i := 0; i < 1000; i++ {
		go CalculateValue(intChantek)
		num := <-intChantek
		fmt.Println(num)
	}

}

const numPool = 1000

func CalculateValue(intChan chan int) {
	random := helpers.RandomNumber(numPool)
	intChan <- random
}
