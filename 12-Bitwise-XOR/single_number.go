package bitwisexor

// SingleNumber finds the only element that appears once in an array
// where all other elements appear exactly twice
// Time Complexity: O(n)
// Space Complexity: O(1)
func SingleNumber(nums []int) int {
	result := 0

	// XOR all numbers in the array
	// Numbers that appear twice will cancel out (a ^ a = 0)
	// The only number that appears once will be left
	for _, num := range nums {
		result ^= num
	}

	return result
}

// SingleNumberIII finds two elements that appear only once in an array
// where all other elements appear exactly twice
// Time Complexity: O(n)
// Space Complexity: O(1)
func SingleNumberIII(nums []int) []int {
	// First, get the XOR of all numbers
	// The result will be the XOR of the two single numbers (a ^ b)
	xorAll := 0
	for _, num := range nums {
		xorAll ^= num
	}

	// Find a bit that is set (1) in xorAll
	// This bit is set in one single number and not in the other
	rightmostSetBit := xorAll & -xorAll // Get the rightmost set bit

	// Divide all numbers into two groups based on this bit
	// Group 1: numbers with this bit set
	// Group 2: numbers with this bit not set
	// The two single numbers will be in different groups
	num1, num2 := 0, 0

	for _, num := range nums {
		if num&rightmostSetBit != 0 {
			num1 ^= num // Number has the bit set
		} else {
			num2 ^= num // Number doesn't have the bit set
		}
	}

	return []int{num1, num2}
}

// SingleNumberII finds the only element that appears once in an array
// where all other elements appear exactly three times
// Time Complexity: O(n)
// Space Complexity: O(1)
func SingleNumberII(nums []int) int {
	// Need to count the bits for each position
	// For each bit position, count the number of times it appears
	// If it's not divisible by 3, that bit is in the single number
	var ones, twos int

	for _, num := range nums {
		// Update twos: Add bits that are in ones AND current number
		twos |= ones & num
		// Update ones: Add current number to ones, remove bits that appear twice
		ones ^= num
		// Create a mask of bits that appear three times
		threes := ones & twos
		// Remove bits that appear three times from ones and twos
		ones &= ^threes
		twos &= ^threes
	}

	return ones
}

// SingleNumberWithDuplicates finds the only element that appears n times in an array
// where all other elements appear m times (where m is not divisible by n)
// This is a general solution
// Time Complexity: O(n * 32) = O(n) where 32 is the number of bits in an integer
// Space Complexity: O(1)
func SingleNumberWithDuplicates(nums []int, n, m int) int {
	// Count bits in each of the 32 positions
	bitCount := make([]int, 32)

	// Count the number of 1s for each bit position
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if (num>>i)&1 == 1 {
				bitCount[i]++
			}
		}
	}

	// Calculate the result by checking each bit position
	result := 0
	for i := 0; i < 32; i++ {
		if bitCount[i]%m != 0 {
			result |= 1 << i
		}
	}

	return result
}

// Example usage:
//
// Single Number:
// Input: nums = [2, 2, 1]
// Output: 1 (1 appears once, all other elements appear twice)
//
// Input: nums = [4, 1, 2, 1, 2]
// Output: 4 (4 appears once, all other elements appear twice)
//
// Single Number III:
// Input: nums = [1, 2, 1, 3, 2, 5]
// Output: [3, 5] (3 and 5 appear once, all other elements appear twice)
//
// Single Number II:
// Input: nums = [2, 2, 3, 2]
// Output: 3 (3 appears once, all other elements appear three times)
//
// Input: nums = [0, 1, 0, 1, 0, 1, 99]
// Output: 99 (99 appears once, all other elements appear three times)
