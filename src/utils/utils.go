package utils

import (
	"fmt"
	"math"
)

// Add возвращает сумму двух целых чисел.
func Add(a, b int) int {
	return a + b
}

// Subtract возвращает разность двух целых чисел.
func Subtract(a, b int) int {
	return a - b
}

// Multiply возвращает произведение двух целых чисел.
func Multiply(a, b int) int {
	return a * b
}

// Divide возвращает результат деления a на b.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

// Power возвращает результат возведения числа a в степень b.
func Power(a float64, b float64) float64 {
	return math.Pow(a, b)
}
