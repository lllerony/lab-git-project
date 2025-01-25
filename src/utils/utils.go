package utils

import (
	"fmt"
	"math"
)

func Add(a, b int) int {
	return a + b // Добавляем +2 для конфликта
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func Power(a float64, b float64) float64 {
	return math.Pow(a, b)
}
