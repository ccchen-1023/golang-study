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
}
