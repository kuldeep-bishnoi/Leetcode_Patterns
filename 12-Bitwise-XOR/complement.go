package bitwisexor

// FindComplement finds the complement of a base 10 number
// The complement is defined as the number formed by flipping all bits of the original number
// We only flip the significant bits (bits used to represent the number)
// Time Complexity: O(log n) where log n is the number of bits in n
// Space Complexity: O(1)
func FindComplement(num int) int {
	if num == 0 {
		return 1 // Edge case: complement of 0 is 1
	}

	// Find the number of bits required to represent num
	// We need this to create a mask with all bits set
	bitCount := 0
	n := num
	for n > 0 {
		bitCount++
		n >>= 1
	}

	// Create a bit mask of all 1's with the same number of bits as num
	// For example, if num = 5 (101 in binary), mask = 7 (111 in binary)
	mask := (1 << bitCount) - 1

	// XOR with mask to flip all bits
	return num ^ mask
}

// BitwiseComplement finds the complement of a base 10 number using another approach
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BitwiseComplement(num int) int {
	if num == 0 {
		return 1 // Edge case: complement of 0 is 1
	}

	// Find a mask that has all bits set to 1 up to the most significant bit of num
	mask := 1
	for mask <= num {
		mask <<= 1
	}

	// Subtract 1 to get all 1's up to the most significant bit of num
	mask -= 1

	// XOR with the mask to flip all bits
	return num ^ mask
}

// FlipBits flips all bits in a 32-bit integer
// This is different from the number's complement as it flips all 32 bits, not just significant bits
// Time Complexity: O(1)
// Space Complexity: O(1)
func FlipBits(num int) int {
	// In Go, ^ is the bitwise NOT operator when used with a single operand
	// But for the function name clarity, we use XOR with all 1's
	return num ^ 0xFFFFFFFF
}

// FlipSignificantBits flips only the significant bits in an integer
// This is the same as FindComplement but with a different implementation
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FlipSignificantBits(num int) int {
	if num == 0 {
		return 1 // Edge case: complement of 0 is 1
	}

	// Find the highest bit that is set
	highestBit := 0
	temp := num
	for temp > 0 {
		highestBit++
		temp >>= 1
	}

	// Create a mask with all 1's up to highest bit
	mask := (1 << highestBit) - 1

	// XOR with the mask
	return num ^ mask
}

// Example usage:
//
// Find Complement:
// Input: num = 5
// Output: 2 (Binary of 5 is 101, its complement is 010, which is 2 in decimal)
//
// Input: num = 1
// Output: 0 (Binary of 1 is 1, its complement is 0)
//
// Flip Bits:
// Input: num = 5
// Output: -6 (In 2's complement, flipping all 32 bits of 5 results in -6)
//
// Flip Significant Bits:
// Input: num = 10
// Output: 5 (Binary of 10 is 1010, flipping significant bits gives 0101, which is 5)
