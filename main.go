package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world."
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	//the function will not print anything to screen until we add println despite the value has been added to the function
	whatWasSaid := SaySomething1()
	fmt.Println("the function returned", whatWasSaid)

	//we can call the function straight away and print some string, because we have been add println on the function
	SaySomething2()

	//function that return 2 value
	whatWasSaid1, theOtherThingToSay := SaySomething3()
	fmt.Println("This fucntion returned", whatWasSaid1, theOtherThingToSay)
}

func SaySomething1() string {
	return "Something"
}

func SaySomething2() string {
	//return "Something"
	whatWasSaid := "something"
	fmt.Println(whatWasSaid)
	return whatWasSaid
}

func SaySomething3() (string, string) {
	return "Something", "else"
}
