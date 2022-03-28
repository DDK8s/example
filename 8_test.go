package main

import (

	"testing"
)

func TestTestingFor(t *testing.T) {
	// Arrange.
	singleList := singleList{}
	expected := []int{1, 2, 3}

	// Act.
	for v := range expected{
		singleList.Push(expected[v])

	}
	result := singleList.TestingFor()

	// Assert.
	for i := range expected {
		if expected[i] != result[i]{

			t.Errorf("Incorrect result. Expected %d, got %d", expected[i], result[i])
			return

		}
	}
}