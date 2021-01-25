package main

import "fmt"

var y = 42

//declare the variable with ID z is of type string and assigned a value "shaken not stirred"
// static programming language, not dynamic
// a variable type cannot be modified
var z = "shaken, not stirred"
var g = "James said \"shaken, not stirred\""

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(g)
	fmt.Printf("%T\n", g)

	// not allowed since z was declared originally string
	//z = 43
	//fmt.Println(z)
	//fmt.Printf("%T\n", z)
}
