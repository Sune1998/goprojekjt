package main

import "fmt"



func main(){
	fmt.Println("enter a number in F")
	var input float64
	fmt.Scanf("%f" , &input)

	output := (input -32) * 5/9

	fmt.Println(output)

	if output <=10 {
		fmt.Println("That's cold my dude")
	} else {
		fmt.Println("very hot")
	}
	print("hej")
}