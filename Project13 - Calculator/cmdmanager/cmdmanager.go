package cmdmanager

import "fmt"

type CMDManager struct {
}

func (fm CMDManager) ReadData() ([]string, error) {
	fmt.Println("Please enter your prices. One price per input. Confirm with Enter. Enter * to finish.")
	var prices []string

	for {
		var price string
		fmt.Print("Price:")
		_, err := fmt.Scan(&price)

		if err != nil {
			return nil, err
		}

		if price == "*" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (fm CMDManager) WriteData(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
