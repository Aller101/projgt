package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSumTable(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{10, 10, 20},
	}

	for _, tt := range tests {
		result := Sum(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Sum(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
