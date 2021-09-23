package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello World", math.Sqrt((4)))
	fmt.Println(rand.Intn(100))
	fmt.Println(add(10, 10))

	var num1 float64 = 1.1 // anything outside of function would need var
	num3 := 111.1

	fmt.Println(add(num1, num3))

	fmt.Println(multiple("Hello", "there"))

	add(num1, num3)

	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(x)
}

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

// Capitalized functiosn are external(imported)

// main will always run

// & tells you memory address
// * gives you value at memory address
