package main

import "fmt"

func main() {
	fmt.Println("hello everyone")
	foo()
	fmt.Println("something else")
	iterative()
}

func iterative() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
