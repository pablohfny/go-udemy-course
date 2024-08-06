package main

import "fmt"

type TransformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)
	double := createTransformer(2)
	tripple := createTransformer(3)
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	trippled := transformNumbers(&numbers, tripple)
	fmt.Println(trippled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transformer TransformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformer(val))
	}
	return dNumbers
}
