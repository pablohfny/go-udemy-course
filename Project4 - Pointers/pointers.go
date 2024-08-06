package main

import "fmt"

func main() {
	// withoutPointers()
	withPointers()
}

// func withoutPointers() {
// 	age := 32
// 	fmt.Println("Age:", age)

// 	adultYears := getAdultYearsWithoutPointers(age)

// 	fmt.Println(adultYears)
// }

// func getAdultYearsWithoutPointers(age int) int {
// 	return age - 18
// }

func withPointers() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age adress:", agePointer)
	fmt.Println("Age value:", *agePointer)

	getAdultYearsWithPointers(&age)

	fmt.Println("Adult Years:", age)
}

func getAdultYearsWithPointers(age *int) {
	*age = *age - 18
}
