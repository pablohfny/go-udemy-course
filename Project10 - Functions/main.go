package main

import "fmt"

type TransformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	doubledReturnFunc := transformNumber(&numbers, getTransformerFunction())
	fmt.Println(doubled)
	fmt.Println(doubledReturnFunc)
	trippled := transformNumber(&numbers, triple)
	fmt.Println(trippled)
	//anonymous function
	quadrupled := transformNumber(&numbers, func(number int) int { return number * 2 })
	fmt.Println(quadrupled)
}

func getTransformerFunction() TransformFn {
	return double
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func transformNumber(numbers *[]int, transformer TransformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformer(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
