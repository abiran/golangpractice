package main

import "fmt"

// declare a variable with an ID > A and type int with value 0 as default
var A int
var z = 44

func main() {
	//short declaration operator
	//declare a variable and assign a value (of a certain type)
	x := 42
	fmt.Println(x)
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	varfoo()
	fmt.Println(A)
	A = 100
	fmt.Println(A)
}

func varfoo() {
	fmt.Println("again:", z)
}
