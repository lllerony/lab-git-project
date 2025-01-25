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
