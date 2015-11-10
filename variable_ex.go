package main

import "fmt"

func main() {
	var x float64
	x = 20.0
	fmt.Println(x)
	// %T is a type format
	/* see more https://golang.org/pkg/fmt/ */
	fmt.Printf("x is of type %T\n", x)
	dynamicType()
	anotherDefind()
	constVariable()
}

func dynamicType() {
	var x float64 = 20.0
	
	y := false
	// z := "this is a wth this \
	// string"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}

func anotherDefind() {
	var a, b, c = 3, true, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}

func constVariable() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("Value of area : %d\n", area)
}