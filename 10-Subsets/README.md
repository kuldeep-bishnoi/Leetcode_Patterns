# Subsets Pattern

## Introduction

The Subsets pattern is used to deal with problems that involve finding permutations, combinations, or subsets of a given set of elements. This pattern uses a systematic approach to build up all possible combinations by incorporating one element at a time, often using techniques like Breadth-First Search (BFS), Depth-First Search (DFS), or Backtracking.

This pattern is particularly useful when you need to explore all possible configurations, scenarios, or combinations of elements that satisfy certain constraints.

## How It Works

The Subsets pattern typically follows these approaches:

1. **Iterative Approach (BFS)**:
   - Start with an empty set
   - For each element, add it to all existing subsets to create new subsets
   - This approach builds the solution one level at a time

2. **Recursive Approach (DFS/Backtracking)**:
   - Make decisions for each element (include it or exclude it)
   - Recursively build subsets by exploring both possibilities
   - Use backtracking to undo choices and explore alternative paths

## Time and Space Complexity

- **Time Complexity**: Often O(N * 2^N) for generating all subsets, where N is the number of elements
- **Space Complexity**: O(2^N) for storing all generated subsets

## When to Use Subsets Pattern

Use the Subsets pattern when:
- You need to find all possible combinations, permutations, or subsets
- You're working with problems that involve exploring different configurations
- You need to generate all possible scenarios or configurations
- You're dealing with problems where order matters (permutations) or doesn't matter (combinations)

## Common Problem Patterns

1. **Subsets/Powerset**
   - Find all possible subsets of a given set

2. **Subsets with Duplicates**
   - Find all possible subsets when the original set contains duplicate elements

3. **Permutations**
   - Generate all possible arrangements (ordering) of a given set

4. **Combinations**
   - Generate all possible ways to select k elements from a set of n elements

5. **Combination Sum**
   - Find all possible combinations of elements that sum to a target value

6. **String Permutations by changing case**
   - Generate all permutations of a string by changing the case of its letters

## Implementation in Golang

```go
// Iterative approach to generate all subsets (powerset)
func Subsets(nums []int) [][]int {
    result := [][]int{{}} // Start with an empty set
    
    for _, num := range nums {
        // For each existing subset, create a new subset by adding the current number
        n := len(result)
        for i := 0; i < n; i++ {
            // Create a new subset by copying the existing subset and adding the current number
            newSet := make([]int, len(result[i]))
            copy(newSet, result[i])
            newSet = append(newSet, num)
            result = append(result, newSet)
        }
    }
    
    return result
}

// Recursive approach (Backtracking) to generate all subsets
func SubsetsRecursive(nums []int) [][]int {
    result := [][]int{}
    currentSet := []int{}
    generateSubsets(nums, 0, currentSet, &result)
    return result
}

func generateSubsets(nums []int, index int, currentSet []int, result *[][]int) {
    // Add the current subset to the result
    subset := make([]int, len(currentSet))
    copy(subset, currentSet)
    *result = append(*result, subset)
    
    // Generate subsets by including elements from the current index onwards
    for i := index; i < len(nums); i++ {
        // Include the current element
        currentSet = append(currentSet, nums[i])
        
        // Recursively generate subsets with the current element included
        generateSubsets(nums, i+1, currentSet, result)
        
        // Backtrack by removing the current element
        currentSet = currentSet[:len(currentSet)-1]
    }
}
```

## Example Problems

1. **Find all Subsets of a Set**
2. **Subsets With Duplicates**
3. **Permutations**
4. **String Permutations by changing case**
5. **Balanced Parentheses**
6. **Unique Generalized Abbreviations**

Each of these problems has a dedicated solution file in this directory. 