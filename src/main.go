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

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
}
