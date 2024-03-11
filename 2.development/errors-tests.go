package main

import (
	"errors"
	"testing"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	// Test case 1: Valid division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result to be 5, got %d", result)
	}

	// Test case 2: Division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected division by zero error, got no error")
	}
}
