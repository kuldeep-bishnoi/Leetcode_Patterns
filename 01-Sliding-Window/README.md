# Sliding Window Pattern

## Introduction

The Sliding Window pattern is a technique used to solve problems involving arrays or lists by maintaining a "window" of elements. This window can grow or shrink as needed, and it "slides" through the array or list to process elements in sequence.

This pattern is particularly useful for solving problems that involve finding subarrays or sublists that satisfy certain conditions, such as:
- Finding the maximum/minimum sum of a subarray of fixed size k
- Finding the longest substring with k distinct characters
- Finding the smallest subarray with a sum greater than or equal to a given value

## How It Works

1. Create a window with either fixed size or variable size depending on the problem
2. Slide the window across the array or list one element at a time
3. When sliding the window, add the new element to the window and remove the element that falls out of the window
4. Process the elements in the current window based on the problem requirements

## Types of Sliding Window

1. **Fixed Size Window**: The window size remains constant throughout the process
2. **Variable Size Window**: The window size can grow or shrink based on the problem constraints

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the size of the array/list
- **Space Complexity**: O(1) to O(k) - where k is the window size or number of unique elements in the window

## When to Use Sliding Window

Use the Sliding Window pattern when:
- You need to process contiguous subarrays or sublists
- The problem involves finding a subrange that meets certain conditions
- You're looking for a maximum, minimum, or optimal result from a subarray

## Common Problem Patterns

1. **Maximum/Minimum Sum Subarray of Size K**
   - Fixed size window problem
   - Find the subarray of size k with the maximum/minimum sum

2. **Longest/Shortest Subarray with a Condition**
   - Variable size window problem
   - Find the longest/shortest subarray that satisfies a condition

3. **String Problems with Constraints**
   - Variable size window problem
   - Find substring with distinct characters, anagrams, etc.

## Implementation in Golang

```go
// Fixed Size Sliding Window - Maximum Sum Subarray of Size K
func maxSumSubarrayOfSizeK(arr []int, k int) int {
    n := len(arr)
    if n < k {
        return -1 // Invalid case
    }
    
    // Compute sum of first window
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    
    maxSum := windowSum
    
    // Slide the window from index k to n
    for i := k; i < n; i++ {
        // Add current element and remove first element of previous window
        windowSum = windowSum + arr[i] - arr[i-k]
        
        // Update maximum sum if required
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    
    return maxSum
}

// Variable Size Sliding Window - Smallest Subarray with Sum >= S
func smallestSubarrayWithGivenSum(arr []int, s int) int {
    n := len(arr)
    windowSum := 0
    minLength := n + 1 // Initialize with a value larger than the array length
    windowStart := 0
    
    for windowEnd := 0; windowEnd < n; windowEnd++ {
        // Add the current element to the window sum
        windowSum += arr[windowEnd]
        
        // Shrink the window as small as possible until the windowSum is less than s
        for windowSum >= s {
            // Update the minimum length
            currentLength := windowEnd - windowStart + 1
            if currentLength < minLength {
                minLength = currentLength
            }
            
            // Remove the leftmost element from the window
            windowSum -= arr[windowStart]
            windowStart++
        }
    }
    
    if minLength > n {
        return 0 // No subarray found
    }
    return minLength
}
```

## Example Problems

1. **Maximum Sum Subarray of Size K**
2. **Longest Substring with K Distinct Characters**
3. **Fruits into Baskets**
4. **Longest Substring Without Repeating Characters**
5. **Permutation in a String**
6. **String Anagrams**
7. **Smallest Window containing Substring**

Each of these problems has a dedicated solution file in this directory. 