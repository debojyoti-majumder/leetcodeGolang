package main

import (
	"testing"
)

func TestNumIdenticalPairs(test *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	ret := numIdenticalPairs(nums)

	if ret != 4 {
		test.Errorf("Expected 4, got %v", ret)
	}

	nums = []int{1, 2, 3}
	ret = numIdenticalPairs(nums)

	if ret != 0 {
		test.Error("Should have returned zero")
	}
}
