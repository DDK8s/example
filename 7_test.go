package main

import (
	"testing"
)

func TestAct(t *testing.T) {
	// Arrange.
	singleList := singleList{}
	var ok bool
	expected := "C"
	var result interface{}


	// Act.
	singleList.fill()
	singleList.makeLoop()
	result, ok = singleList.findLoop()

	// Assert.
	if !ok && expected != result {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, result)
		return
	}

}
