package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := sumup(numbers)
	sum2 := variadicSumup(1, 2, 3, 4)
	sum3 := variadicSumup(numbers...)
	fmt.Println(sum)
	fmt.Println(sum2)
	fmt.Println(sum3)
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func variadicSumup(numbers ...int) int {
	return sumup(numbers)
}
