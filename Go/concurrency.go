package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {
	defer fmt.Println("Finsish defer") // will only run if the rest of the foo func is done/ or error/panic out
	defer fmt.Println("Finsish defer2")
	fmt.Println("SOmething")
}

func cleanup() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered in cleanup", r) // r will be the panic
	}

}

func say(s string) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("panicccccs")
		}
	}
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
	// Above does not print anything if waitgroup is not implemented before the program ends before the conccurrent threads ran
	//time.Sleep(time.Second)
}
