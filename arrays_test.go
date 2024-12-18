package main

import (
	"reflect"
	"testing"
)

func TestArrays(t *testing.T) {
	// Test array initialization and element access
	var a [5]int
	if a[4] != 0 {
		t.Errorf("Expected a[4] to be 0, got %d", a[4])
	}
	a[4] = 100
	if a[4] != 100 {
		t.Errorf("Expected a[4] to be 100, got %d", a[4])
	}

	// Test array length
	if len(a) != 5 {
		t.Errorf("Expected length of a to be 5, got %d", len(a))
	}

	// Test array literal declaration
	b := [5]int{1, 2, 3, 4, 5}
	expectedB := [5]int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(b, expectedB) {
		t.Errorf("Expected b to be %v, got %v", expectedB, b)
	}

	// Test array literal with ellipsis
	b = [...]int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(b, expectedB) {
		t.Errorf("Expected b to be %v, got %v", expectedB, b)
	}

	// Test array literal with index and value pairs
	b = [...]int{100, 3: 400, 500}
	expectedB = [5]int{100, 0, 0, 400, 500}
	if !reflect.DeepEqual(b, expectedB) {
		t.Errorf("Expected b to be %v, got %v", expectedB, b)
	}

	// Test 2D array initialization
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	expected2D := [2][3]int{{0, 1, 2}, {1, 2, 3}}
	if !reflect.DeepEqual(twoD, expected2D) {
		t.Errorf("Expected twoD to be %v, got %v", expected2D, twoD)
	}

	// Test 2D array literal declaration
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	expected2D = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	if !reflect.DeepEqual(twoD, expected2D) {
		t.Errorf("Expected twoD to be %v, got %v", expected2D, twoD)
	}
}
