package main

import "fmt"

func main() {
	fmt.Println("enter a number in F") 			// dette program konvertere F til C og fortæller dig om det er koldt eller varmt.
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

	patrick()
}

//hej
// hej sune, hilsen patrick
func patrick() {        				// denne funktion spørg om dit navn 
	var input string

	fmt.Println("Hello, World! This is my first GO program")
	fmt.Print("What's your name? ")
	fmt.Scanf("%s", &input)
	output := "Greetings, you are my new friend " + input

	fmt.Println(output)

	var (
		A = 10
		B = 17
		C = 7
	)
	fmt.Println("My name is Patrick and i am turning ", A+B, " years old this month")
	fmt.Println("I miss being ", (A+B)-C, " years old")

}
