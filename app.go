package main

import (
	"fmt"
	"strconv"
)

//Global -- PascalCase

// var val2 int = 100

// package level -- camelCase
// var myNum int = 88

func main() {
	//var a int = 48 // Declaretion
	// fmt.Print(a)
	// fmt.Println(a)
	// fmt.Println(a)
	// fmt.Println(a)

	var a1 bool = true    // Boolean
	var b1 int = 5        // Integer
	var c1 float32 = 3.14 // Floating point number
	var d1 string = "Hi!" // String

	fmt.Println("Boolean: ", a1)
	fmt.Println("Integer: ", b1)
	fmt.Println("Float:   ", c1)
	fmt.Println("String:  ", d1)

	// b := 4488
	// fmt.Println(b)
	// Local
	// var val int = 48
	// fmt.Printf("%v, %T", val, val2)

	var a int = 48
	var b string = strconv.Itoa(a)
	fmt.Printf("%v, %T", b, b)
}
