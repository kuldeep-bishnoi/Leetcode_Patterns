# Two Pointers Pattern

## Introduction

The Two Pointers pattern is a technique that involves using two pointers to iterate through a data structure. These pointers can move toward each other (from both ends), move in the same direction at different speeds, or follow other patterns depending on the problem.

This pattern is particularly effective for solving problems involving sorted arrays, linked lists, or situations where we need to find pairs of elements that satisfy certain conditions.

## How It Works

1. Initialize two pointers to specific positions in the data structure
2. Move the pointers based on the problem's requirements until they meet a condition
3. Process the elements at the pointers' positions

## Types of Two Pointers

1. **Opposite Direction (Left and Right)**: Pointers start at opposite ends and move toward each other
2. **Same Direction (Fast and Slow)**: Both pointers move in the same direction but at different speeds
3. **Same Direction (Window)**: One pointer waits while the other moves, creating a window

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the size of the array or list
- **Space Complexity**: O(1) - as we're usually operating in place

## When to Use Two Pointers

Use the Two Pointers pattern when:
- You're working with a sorted array
- You need to find pairs of elements that satisfy certain conditions
- You're dealing with palindromes or symmetry
- You're searching for a target sum
- You need to remove duplicates in place

## Common Problem Patterns

1. **Pair with Target Sum**
   - Find a pair of elements in a sorted array that sum to a target value
   - Start pointers at opposite ends and move them toward each other

2. **Remove Duplicates**
   - Remove duplicate elements from a sorted array in place
   - Use two pointers moving in the same direction at different paces

3. **Squaring a Sorted Array**
   - Square a sorted array of integers while maintaining the sorted order
   - Compare absolute values from both ends and place squares in a new array

4. **Triplet Sum to Zero (3Sum)**
   - Find all triplets in an array that sum to zero
   - Sort the array, fix one element, and use two pointers to find the other two

## Implementation in Golang

```go
// Pair with Target Sum
func pairWithTargetSum(arr []int, targetSum int) []int {
    left, right := 0, len(arr)-1
    
    for left < right {
        currentSum := arr[left] + arr[right]
        
        if currentSum == targetSum {
            return []int{left, right} // Found the pair
        }
        
        if currentSum < targetSum {
            left++ // Need a pair with a larger sum
        } else {
            right-- // Need a pair with a smaller sum
        }
    }
    
    return []int{-1, -1} // No pair found
}

// Remove Duplicates
func removeDuplicates(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    
    // 'nextNonDuplicate' is the index where the next non-duplicate element should go
    nextNonDuplicate := 1
    
    for i := 1; i < len(arr); i++ {
        if arr[nextNonDuplicate-1] != arr[i] {
            arr[nextNonDuplicate] = arr[i]
            nextNonDuplicate++
        }
    }
    
    return nextNonDuplicate // Length of the array after removing duplicates
}

// Squaring a Sorted Array
func sortedArraySquares(arr []int) []int {
    n := len(arr)
    squares := make([]int, n)
    highestSquareIdx := n - 1
    left, right := 0, n-1
    
    for left <= right {
        leftSquare := arr[left] * arr[left]
        rightSquare := arr[right] * arr[right]
        
        if leftSquare > rightSquare {
            squares[highestSquareIdx] = leftSquare
            left++
        } else {
            squares[highestSquareIdx] = rightSquare
            right--
        }
        
        highestSquareIdx--
    }
    
    return squares
}
```

## Example Problems

1. **Pair with Target Sum**
2. **Remove Duplicates**
3. **Squaring a Sorted Array**
4. **Triplet Sum to Zero (3Sum)**
5. **Triplet Sum Close to Target**
6. **Triplets with Smaller Sum**
7. **Subarray Product Less Than K**
8. **Dutch National Flag Problem (Sort Colors)**
9. **4Sum**
10. **Backspace String Compare**
11. **Shortest Unsorted Continuous Subarray**

Each of these problems has a dedicated solution file in this directory. 