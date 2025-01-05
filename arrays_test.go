package main

import (
	"reflect"
	"testing"
)

func TestSetArray(t *testing.T) {
	t.Run("in range", func(t *testing.T) {
		var a [10]int
		expected := [10]int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		got := setArray(a[:], 0, 100)

		if !reflect.DeepEqual(a, expected) {
			t.Errorf("Expected %v, got %v", expected, a)
		}
		if !got {
			t.Errorf("Expected true, got %v", got)
		}
	})

	t.Run("out of range", func(t *testing.T) {
		var a [10]int
		got := setArray(a[:], 15, 100)
		if got {
			t.Errorf("Expected false, got %v", got)
		}
	})

}
