package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age int
}

func main() {
	//instantiate variabel of type person (struct)
	p := person{name: "peter", age: 10}
	fmt.Println(p)

	//Map
	vertices := make(map[string]int)
	vertices["Triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12



	//Variable examples
	x := 5
	y := 2
	xysum := x+y
	result := sum(2, 5)
	fmt.Println(result)

	//array
	b := [5]int{2,1,3,1,2}

	//slice is an array with an undefined length, like arraylist
	a := []int{2,2,2,2,1,3}

	//add element to slice
	a = append(a, 1123123)

	//Console print
	fmt.Println(vertices)
	fmt.Println(a)
	fmt.Println(b)


	//If statement
	if x > y {
		fmt.Println(":)")
	} else {
		fmt.Println(xysum)
		fmt.Println("Hello World")
	}

	//There is only for loops in go.
	for i:=0; i < 5; i++ {
		fmt.Print(i)
	}

	//You can loop through arrays
	for index, value := range a {
		fmt.Println("index: ", index, "value: ", value)
		
	}

	//You can loop through maps
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	result2, err := sqrt(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

}

//Method example where return type is placed after ().
func sum(x int, y int) int {
	return x+y
}

//
func sqrt(x float64) (float64, error)  {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers.")
	} else  {
		return math.Sqrt(x), nil
	}
}