package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount float64 = 1000
	// expectedReturnRate := 5.5
	// var years float64 = 10
	const inflationRate = 2.5
	var investmentAmount, returnRate, years float64

	fmt.Print("Investment Amount ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Return Rate ")
	fmt.Scan(&returnRate)

	fmt.Print("Years ")
	fmt.Scan(&years)

	futureValue := calculateFutureValue(investmentAmount, returnRate, years)
	futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)

	fmt.Printf("Future Value: %.2f\nFuture Value (Adjusted for Inflation): %.2f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, returnRate float64, years float64) float64 {
	return investmentAmount * math.Pow((1+returnRate/100), years)
}

func calculateFutureRealValue(futureValue float64, inflationRate float64, years float64) float64 {
	return futureValue / (math.Pow(1+inflationRate/100, years))
}
