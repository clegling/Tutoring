package tests

import (
	"Tutorial/tasks"
	"testing"
)

func TestDuplicates(t *testing.T) {
	var result bool
	result = tasks.CheckDuplicates([]int{1, 2, 3, 4, 5, 6, 7})
	if result {
		t.Errorf("Error for Task 1")
	}
	result = tasks.CheckDuplicates([]int{1, 2, 3, 4, 5, 6, 6})
	if !result {
		t.Errorf("Error for Task 2")
	}

}
