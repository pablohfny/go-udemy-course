package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"google": "https://google.com",
		"aws":    "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "google")
	fmt.Println(websites)

	//Space preallocation
	userNames := make([]string, 2)
	fmt.Println(userNames)
	userNames = append(userNames, "1")
	userNames = append(userNames, "2")
	fmt.Println(userNames)
	userNames[0] = "3"
	userNames[1] = "2"
	fmt.Println(userNames)

	passwords := make([]string, 0, 5)
	fmt.Println(passwords)
	passwords = append(passwords, "p1")
	passwords = append(passwords, "p1")
	passwords = append(passwords, "p1")
	passwords = append(passwords, "p1")
	passwords = append(passwords, "p1")
	fmt.Println(passwords)

	//Making maps
	courseRatings := make(floatMap, 2)
	courseRatings["teste"] = 1
	courseRatings["teste2"] = 2
	courseRatings.output()

	//Iterating over maps
	for index, value := range courseRatings {
		fmt.Println("key:", index)
		fmt.Println("value:", value)
	}
}
