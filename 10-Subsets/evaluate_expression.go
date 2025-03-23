package subsets

import (
	"strconv"
)

// DiffWaysToCompute computes all possible results from an arithmetic expression
// by adding different parenthesizations
// Time Complexity: O(4^n) in the worst case, where n is the length of the expression
// Space Complexity: O(4^n) for storing all the results
func DiffWaysToCompute(expression string) []int {
	// Memoization map to avoid redundant calculations
	memo := make(map[string][]int)
	return diffWaysToComputeHelper(expression, memo)
}

// diffWaysToComputeHelper recursively evaluates different ways to compute the expression
func diffWaysToComputeHelper(expression string, memo map[string][]int) []int {
	// Check if we've already computed this expression
	if result, exists := memo[expression]; exists {
		return result
	}

	results := []int{}

	// Try to parse the entire expression as a single number
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}

	// Iterate through each character of the expression
	for i := 0; i < len(expression); i++ {
		char := expression[i]

		// If we find an operator, split the expression and evaluate both sides
		if char == '+' || char == '-' || char == '*' {
			// Compute all results from the left part
			leftResults := diffWaysToComputeHelper(expression[:i], memo)

			// Compute all results from the right part
			rightResults := diffWaysToComputeHelper(expression[i+1:], memo)

			// Combine left and right results with the current operator
			for _, left := range leftResults {
				for _, right := range rightResults {
					var result int

					switch char {
					case '+':
						result = left + right
					case '-':
						result = left - right
					case '*':
						result = left * right
					}

					results = append(results, result)
				}
			}
		}
	}

	// Cache the results
	memo[expression] = results
	return results
}

// AddOperators finds all possible ways to add operators (+, -, *) to a digit string
// such that the resulting expression evaluates to the target value
// Time Complexity: O(4^n) where n is the length of the digit string
// Space Complexity: O(n) for recursion stack and storing valid expressions
func AddOperators(num string, target int) []string {
	result := []string{}
	if len(num) == 0 {
		return result
	}

	backtrackAddOperators(num, target, 0, 0, 0, "", &result)
	return result
}

// backtrackAddOperators recursively explores different ways to add operators
// index: current position in the num string
// eval: current evaluation result
// multed: the multiplier in the current context (for handling * operator)
// path: current expression path
func backtrackAddOperators(num string, target int, index int, eval, multed int64, path string, result *[]string) {
	// If we've processed the entire string
	if index == len(num) {
		// If the evaluation matches the target, add to result
		if eval == int64(target) {
			*result = append(*result, path)
		}
		return
	}

	// Try different lengths of the current number
	for i := index; i < len(num); i++ {
		// Skip leading zeros
		if i > index && num[index] == '0' {
			break
		}

		// Parse current number
		currStr := num[index : i+1]
		curr, _ := strconv.ParseInt(currStr, 10, 64)

		// For the first number, no operator needed
		if index == 0 {
			backtrackAddOperators(num, target, i+1, curr, curr, currStr, result)
			continue
		}

		// Add "+" operator
		backtrackAddOperators(num, target, i+1, eval+curr, curr, path+"+"+currStr, result)

		// Add "-" operator
		backtrackAddOperators(num, target, i+1, eval-curr, -curr, path+"-"+currStr, result)

		// Add "*" operator (needs to handle precedence)
		// eval - multed + (multed * curr) accounts for the precedence of multiplication
		backtrackAddOperators(num, target, i+1, eval-multed+(multed*curr), multed*curr, path+"*"+currStr, result)
	}
}

// Example usage:
//
// For DiffWaysToCompute:
// Input: expression = "2-1-1"
// Output: [0, 2]
// Explanation:
// ((2-1)-1) = 0
// (2-(1-1)) = 2
//
// For AddOperators:
// Input: num = "123", target = 6
// Output: ["1+2+3", "1*2*3"]
// Explanation: Both "1+2+3" and "1*2*3" evaluate to 6.
