package main

import "fmt"

//	func main() {
//		var productNames = [2]string{"a"}
//		fmt.Println(productNames)
//		fmt.Println(productNames[:1])
//		fmt.Println(len(productNames), cap(productNames))
//	}
func main() {
	prices := []float64{}
	prices2 := []float64{1.0, 2.0}
	fmt.Println(prices)
	prices = append(prices, 1.1)
	fmt.Println(prices)
	prices = append(prices, prices2...)
	fmt.Appendln(prices)
}
