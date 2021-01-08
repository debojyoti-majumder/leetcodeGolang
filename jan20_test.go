package main

import (
	"reflect"
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

func TestRunningSum(test *testing.T) {
	nums := []int{1, 2, 3, 4}
	ret := runningSum(nums)
	expectedOutcome := []int{1, 3, 6, 10}

	if len(ret) != len(nums) {
		test.Error("Length mismatch")
	}

	if reflect.DeepEqual(ret, expectedOutcome) == false {
		test.Errorf("Retruned %v", ret)
	}

	nums = []int{}
	ret = runningSum(nums)

	if len(ret) != 0 {
		test.Error("Non zero slice")
	}
}
