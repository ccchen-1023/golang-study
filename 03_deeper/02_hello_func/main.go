package main

import (
	"fmt"
)

func main() {
	greeting := newGreeting()

	fmt.Println(greeting)

	// go build main.go
	// ./main

	print("aloha")
}

func newGreeting() string {
	return "hello"
}

func print(greeting string) {
	fmt.Println(greeting)
}
