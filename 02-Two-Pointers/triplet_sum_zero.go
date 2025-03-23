package twopointers

import (
	"sort"
)

// TripletSumToZero finds all unique triplets in the array that sum up to zero
// Time Complexity: O(n²) where n is the length of the array
// Space Complexity: O(n) for the sorting
func TripletSumToZero(arr []int) [][]int {
	result := [][]int{}
	if len(arr) < 3 {
		return result
	}

	// Sort the array
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		// Skip duplicate triplets
		// If the current number is the same as the previous one, skip it
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		// Find pairs with a sum of -arr[i]
		targetSum := -arr[i]
		left := i + 1
		right := len(arr) - 1

		for left < right {
			currentSum := arr[left] + arr[right]

			if currentSum == targetSum { // Found a triplet
				result = append(result, []int{arr[i], arr[left], arr[right]})

				// Skip duplicate numbers
				left++
				for left < right && arr[left] == arr[left-1] {
					left++
				}

				right--
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			} else if currentSum < targetSum {
				left++ // Need a pair with a larger sum
			} else {
				right-- // Need a pair with a smaller sum
			}
		}
	}

	return result
}

// Example cases:
// Input: [-3, 0, 1, 2, -1, 1, -2]
// Output: [[-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]]
// Explanation: There are four unique triplets whose sum is equal to zero

// Input: [-5, 2, -1, -2, 3]
// Output: [[-5, 2, 3], [-2, -1, 3]]
// Explanation: There are two unique triplets whose sum is equal to zero
