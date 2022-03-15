package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	// Arrange.
	expected := []int{1, 2, 3, 4, 6, 7, 8}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act.
	result := worker(numbers)

	// Assert.

			if result[cap(expected)-1] != expected[cap(expected)-1] {
				t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
				return
			}



}