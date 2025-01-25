package main

import (
	"lab-git-project/src/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	result := utils.Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestSubtract(t *testing.T) {
	result := utils.Subtract(5, 2)
	if result != 3 {
		t.Errorf("Subtract(5, 2) = %d; want 3", result)
	}
}

func TestMultiply(t *testing.T) {
	result := utils.Multiply(4, 2)
	if result != 8 {
		t.Errorf("Multiply(4, 2) = %d; want 8", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := utils.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned an error: %s", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}

	_, err = utils.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return an error")
	}
}
