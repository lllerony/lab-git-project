package main

import (
	"fmt"
	"lab-git-project/src/utils"
)

func main() {
	a := 10
	b := 5

	sum := utils.Add(a, b)
	diff := utils.Subtract(a, b)
	product := utils.Multiply(a, b)
	quotient, err := utils.Divide(a, b)
	power := utils.Power(float64(a), float64(b))

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	if err != nil {
		fmt.Printf("Division error: %s\n", err)
	} else {
		fmt.Printf("Quotient: %d\n", quotient)
	}
	fmt.Printf("Power: %.2f\n", power)
	fmt.Printf("Done!")
}
