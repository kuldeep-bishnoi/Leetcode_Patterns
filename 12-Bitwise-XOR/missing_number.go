package bitwisexor

// FindMissingNumber finds the missing number in an array containing n distinct numbers
// taken from 0 to n (inclusive)
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMissingNumber(nums []int) int {
	n := len(nums)

	// XOR all numbers from 0 to n
	// XOR all numbers in the array
	// The result will be the missing number
	xor := 0

	// XOR all numbers from 0 to n
	for i := 0; i <= n; i++ {
		xor ^= i
	}

	// XOR with all numbers in the array
	for _, num := range nums {
		xor ^= num
	}

	return xor
}

// FindMissingNumberMath finds the missing number using a mathematical formula
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMissingNumberMath(nums []int) int {
	n := len(nums)

	// Sum of numbers from 0 to n can be calculated using the formula n*(n+1)/2
	expectedSum := n * (n + 1) / 2

	// Calculate the actual sum
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	// The difference is the missing number
	return expectedSum - actualSum
}

// FindMissingNumberAlternative finds the missing number using an alternative XOR approach
// by combining both loops
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMissingNumberAlternative(nums []int) int {
	n := len(nums)
	result := n // Start with n (the largest possible value)

	// XOR with index and value simultaneously
	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}

	return result
}

// FindMissingNumbers finds two missing numbers in an array containing n-2 distinct numbers
// taken from 0 to n (inclusive)
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMissingNumbers(nums []int) []int {
	n := len(nums) + 2 // The array has n-2 elements, so n is len(nums) + 2

	// First, get the XOR of all numbers from 0 to n and all numbers in the array
	xorAll := 0

	// XOR all numbers from 0 to n
	for i := 0; i <= n; i++ {
		xorAll ^= i
	}

	// XOR with all numbers in the array
	for _, num := range nums {
		xorAll ^= num
	}

	// Now, xorAll is the XOR of the two missing numbers
	// Find a bit that is set (1) in xorAll
	// This bit is set in one missing number and not in the other
	rightmostSetBit := xorAll & -xorAll // Get the rightmost set bit

	// Divide all numbers into two groups based on this bit
	// Group 1: numbers with this bit set
	// Group 2: numbers with this bit not set
	// The two missing numbers will be in different groups
	group1, group2 := 0, 0

	// Check numbers from 0 to n
	for i := 0; i <= n; i++ {
		if i&rightmostSetBit != 0 {
			group1 ^= i // i has the bit set
		} else {
			group2 ^= i // i doesn't have the bit set
		}
	}

	// Check numbers in the array
	for _, num := range nums {
		if num&rightmostSetBit != 0 {
			group1 ^= num // num has the bit set
		} else {
			group2 ^= num // num doesn't have the bit set
		}
	}

	return []int{group1, group2}
}

// Example usage:
//
// Find Missing Number:
// Input: nums = [3, 0, 1]
// Output: 2 (the missing number is 2)
//
// Input: nums = [9, 6, 4, 2, 3, 5, 7, 0, 1]
// Output: 8 (the missing number is 8)
//
// Find Two Missing Numbers:
// Input: nums = [1, 2, 4]
// Output: [0, 3] (the missing numbers are 0 and 3)
//
// Input: nums = [0, 1, 3, 4, 5]
// Output: [2, 6] (the missing numbers are 2 and 6)
