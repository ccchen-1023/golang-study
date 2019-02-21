package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "2018"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
