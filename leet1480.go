package main

func runningSum(nums []int) []int {
	numberOfElements := len(nums)

	// No need to perform sums as the list is small
	if numberOfElements <= 1 {
		return nums
	}

	// We know the length of the return value in advance so we can create
	// the array
	runningSumValues := make([]int, numberOfElements)
	runningSumValues[0] = nums[0]

	// Calculating the running sum
	for i := 1; i < numberOfElements; i++ {
		runningSumValues[i] = runningSumValues[i-1] + nums[i]
	}

	return runningSumValues
}
