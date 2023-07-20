package main

import "fmt"

func main() {
	fmt.Println(sum(4, 5, 6))
	fmt.Println(mult(4, 5, 6))
	fmt.Println(sub(5, 3))
	fmt.Println(div(4, 2))
}

func sum(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

func mult(values ...int) int {
	total := 1

	for _, v := range values {
		total *= v
	}

	return total
}

func sub(firstValue, secondValue int) int {
	return firstValue - secondValue
}

func div(firstValue, secondValue int) float64 {
	return float64(firstValue) / float64(secondValue)
}
