// executable package
package main

import (
	// reuseable package
	"fmt"
)

func main() {
	var greeting string = "hello"
	// should omit type string from declaration of var greeting; it will be inferred from the right-hand side
	// var greeting := "hello"
	greeting = "hi"

	// go run main.go
	fmt.Println(greeting)

	var i, j int
	fmt.Println(i, j)

	var s, f string
	fmt.Println(s == "", f)

	pi := 3.14
	no := 5
	fmt.Println(pi * float64(no))

	fmt.Println(int(pi))
}
