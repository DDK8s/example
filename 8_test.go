package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	// Arrange.
	var sls []int
	singList := initList()
	var setOfStacks []*element

	expectedSet := []*element{
		{data: 7, next: nil},
		{data: 4, next: nil},
		{data: 7, next: nil},

	}

	numbers := []int{7, 2, 3, 4, 5, 6, 7, 8, 9}
	//expected := 2

	// Act.
	result, sls := singList.Push(setOfStacks, singList, numbers, sls)

	// Assert.
	for v := range result {
		for d := range expectedSet {
			if result[v].data != expectedSet[d].data{
				t.Errorf("Incorrect result. Expected %d, got %d", expectedSet[d].data, result[v].data)
				return
			}
		}
	}

}