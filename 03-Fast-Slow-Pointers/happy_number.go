package fastslow

// IsHappy determines if a number is a "happy number"
// A happy number is a number that eventually reaches 1 when replaced by the sum of the square of each digit
// Time Complexity: O(log n)
// Space Complexity: O(1)
func IsHappy(n int) bool {
	if n <= 0 {
		return false
	}

	// Initialize slow and fast pointers
	slow, fast := n, squareSum(n)

	// If n is a happy number, the sequence will eventually reach 1
	// If not, it will enter a cycle, and fast will eventually meet slow
	for fast != 1 && slow != fast {
		slow = squareSum(slow)            // Move one step
		fast = squareSum(squareSum(fast)) // Move two steps
	}

	// If fast pointer reached 1, it's a happy number
	return fast == 1
}

// squareSum calculates the sum of squares of digits of a number
func squareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// Example cases:
// Input: 19
// Output: true
// Explanation:
// 1² + 9² = 1 + 81 = 82
// 8² + 2² = 64 + 4 = 68
// 6² + 8² = 36 + 64 = 100
// 1² + 0² + 0² = 1 + 0 + 0 = 1 (reaches 1, so it's a happy number)

// Input: 2
// Output: false
// Explanation:
// 2² = 4
// 4² = 16
// 1² + 6² = 1 + 36 = 37
// 3² + 7² = 9 + 49 = 58
// 5² + 8² = 25 + 64 = 89
// 8² + 9² = 64 + 81 = 145
// 1² + 4² + 5² = 1 + 16 + 25 = 42
// 4² + 2² = 16 + 4 = 20
// 2² + 0² = 4 + 0 = 4
// This enters a cycle: 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4
// Since it never reaches 1, it's not a happy number
