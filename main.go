package main

import "fmt"

func main() {
	fmt.Println("enter a number in F")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output)

	if output <= 10 {
		fmt.Println("That's cold my dude")
	} else {
		fmt.Println("very hot")
	}
	println("hej")
	//calling patrick method
	patrick()
}

//method patrick, sets input type to string
func patrick() {
	var input string
	//prints souts
	fmt.Println("Hello, World! This is my first GO program")
	fmt.Print("What's your name? ")
	//starts scanner, takes string
	fmt.Scanf("%s", &input)
	//declares text string + input to output
	output := "Greetings, you are my new friend " + input
	//souts input
	fmt.Println(output)

	//declaring variables
	var (
		A = 10
		B = 17
		C = 7
	)
	//souting text + variables
	fmt.Println("My name is Patrick and i am turning ", A+B, " years old this month")
	fmt.Println("I miss being ", (A+B)-C, " years old")
	//calls method proofOfLoop
	proofOfLoop()
}

//Method with loop, prints "GRUPPE GO-UDVIKLING"
//calls proofOfArray method 5 times.
func proofOfLoop() {
	//loop running 5 times
	for i := 0; i < 5; i++ {
		fmt.Print("GRUPPE GO-UDVIKLING\n")
		//calls method proofOfArray
		proofOfArray()
	}
}

//Method with array, declaring array a,
//filling it with group names and prints (a)
func proofOfArray() {
	//declaring array (a)
	a := []string{"Bro ", "Kelvin ", "Rasmus ", "Patrick "}
	fmt.Print(a)
}
