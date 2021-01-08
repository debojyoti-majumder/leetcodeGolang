// URL : 	https://leetcode.com/problems/number-of-good-pairs/
// Status: 	Accepted

// Related Problems:
//	https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
//  https://leetcode.com/problems/corporate-flight-bookings/

package main

func numIdenticalPairs(nums []int) int {
	pairCount := 0
	numberOfElements := len(nums)

	// Basic check
	if numberOfElements <= 1 {
		return 0
	}

	// Creating the number map which would behave like a set
	numbersMap := map[int]int{nums[0]: 1}

	// Going through the numebrs
	for i := 1; i < numberOfElements; i++ {
		number := nums[i]
		value, ok := numbersMap[number]

		if ok {
			pairCount += value
			numbersMap[number]++
		} else {
			numbersMap[number] = 1
		}
	}

	return pairCount
}
