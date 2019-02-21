package main

import "fmt"

func main() {

	numbers := []int{}

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for i, number := range numbers {
		result := "odd"
		if number%2 == 0 {
			result = "even"
		}
		fmt.Println(i, "is", result)
	}

}
