package bitwisexor

// TwoSingleNumbers finds two elements that appear only once in an array
// where all other elements appear exactly twice
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSingleNumbers(nums []int) []int {
	// First, get the XOR of all numbers in the array
	// The result will be the XOR of the two single numbers (a ^ b)
	xorAll := 0
	for _, num := range nums {
		xorAll ^= num
	}

	// Find a bit that is set (1) in xorAll
	// This bit is different in the two single numbers
	// (one number has this bit set, the other doesn't)
	rightmostSetBit := xorAll & -xorAll // Get the rightmost set bit

	// Divide all numbers into two groups:
	// Group 1: numbers with the bit set
	// Group 2: numbers with the bit not set
	// The two single numbers will be in different groups
	num1, num2 := 0, 0

	for _, num := range nums {
		if num&rightmostSetBit != 0 {
			// Group 1: The bit is set
			num1 ^= num
		} else {
			// Group 2: The bit is not set
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

// TwoSingleNumbersAlternative finds two elements that appear only once
// using a slightly different approach
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSingleNumbersAlternative(nums []int) []int {
	// XOR all numbers to get a ^ b
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// Get any set bit position where a and b differ
	// Here we use a different approach to find a bit that is set
	diffBit := 1
	for (xor & diffBit) == 0 {
		diffBit <<= 1
	}

	// Separate numbers into two groups
	a, b := 0, 0
	for _, num := range nums {
		if num&diffBit == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}

// TwoSingleNumbersWithOrder finds two elements that appear only once
// and returns them in ascending order
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSingleNumbersWithOrder(nums []int) []int {
	result := TwoSingleNumbers(nums)

	// Ensure the result is in ascending order
	if result[0] > result[1] {
		result[0], result[1] = result[1], result[0]
	}

	return result
}

// TwoSingleNumbersWithDifferentMultiplicity finds two elements where
// one element appears 'p' times and all other elements (except one) appear 'q' times
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSingleNumbersWithDifferentMultiplicity(nums []int, p, q int) []int {
	// Count bits for each position (32-bit integer)
	bitCount := make([]int, 32)

	// Count the number of 1s for each bit position
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if (num>>i)&1 == 1 {
				bitCount[i]++
			}
		}
	}

	// The result numbers
	a, b := 0, 0

	// For each bit position, determine which number contains that bit
	for i := 0; i < 32; i++ {
		bit := 1 << i
		remainder := bitCount[i] % q

		if remainder == 0 {
			// This bit position is not part of either special number
			continue
		}

		if remainder == p%q {
			// This bit is part of the number that appears p times
			a |= bit
		} else {
			// This bit is part of the number that appears once
			b |= bit
		}
	}

	return []int{a, b}
}

// Example usage:
//
// Two Single Numbers:
// Input: nums = [1, 2, 1, 3, 2, 5]
// Output: [3, 5] (3 and 5 appear once, all other elements appear twice)
//
// Input: nums = [4, 1, 2, 1, 2]
// Output: [4, 0] (4 appears once, all other elements appear twice; 0 is implied as the second single number)
//
// Two Single Numbers With Order:
// Input: nums = [1, 2, 1, 3, 2, 5]
// Output: [3, 5] (3 and 5 in ascending order)
//
// Two Single Numbers With Different Multiplicity:
// Input: nums = [1, 1, 1, 2, 2, 3], p = 3, q = 2
// Output: [1, 3] (1 appears 3 times, 3 appears once, all others appear twice)
