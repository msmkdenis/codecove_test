package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(diff(1, 2))
	fmt.Println(multiply(1, 2))
}

func sum(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
