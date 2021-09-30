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

	for i := 0; i < 10; i++ { //FOR Loop
		fmt.Println(i)
	}

	i := 0
	for i < 10 { // WHILE Loop
		fmt.Println(i)
		i++
	}

	// for {						  // Infinite Loop
	// 	fmt.Println("INFINITE LOOP")
	// }

	z := 5

	for {
		fmt.Println("Do Stuff", z)
		z += 3
		if z > 25 {
			break
		}
	}

	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jay"] = 95

	TimsGrade := grades["Timmy"]
	fmt.Print(TimsGrade)
	delete(grades, "Timmy")

	for key, value := range grades {
		fmt.Print(key, value)
	}

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
