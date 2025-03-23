# Cyclic Sort Pattern

## Introduction

The Cyclic Sort pattern is an efficient sorting technique specifically designed for problems where the input array contains numbers in a given range. This pattern leverages the fact that we know the positions where each element should be in the final sorted array.

The main idea behind this pattern is to place each number at its correct index in the array. For example, if the input array contains numbers from 1 to N, then each number should be placed at the index that is one less than its value (since array indices start at 0).

## How It Works

1. Iterate through the array and check if the current element is in its correct position
2. If the element is not in the correct position, swap it with the element at its intended position
3. After swapping, stay at the same index (don't increment the index) and check the new element
4. If the element is already in the correct position, move to the next element

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the size of the input array
- **Space Complexity**: O(1) - since we are sorting in-place

## When to Use Cyclic Sort

Use the Cyclic Sort pattern when:
- The problem involves arrays containing numbers in a given range
- The problem asks to find the missing/duplicate/smallest number in an array
- You need to sort the array in-place with minimal swaps

## Common Problem Patterns

1. **Sort an Array with Numbers in Range 1 to n**
   - Place each number at index (number - 1)
   - After sorting, the array should have each number at its correct position

2. **Find the Missing Number**
   - After cyclic sort, the first index where the value doesn't match (index + 1) is the missing number
   - If all elements are in their correct positions, then n is the missing number

3. **Find All Missing Numbers**
   - Similar to finding one missing number, but collect all indices where the value doesn't match

4. **Find the Duplicate Number**
   - If during swapping we find an element already at its correct position, that's the duplicate

5. **Find All Duplicate Numbers**
   - Track all elements that are already at their correct positions during the sort

## Implementation in Golang

```go
// Cyclic Sort for array with numbers 1 to n
func cyclicSort(nums []int) {
    i := 0
    for i < len(nums) {
        // Correct position for the current number
        correctPos := nums[i] - 1
        
        // If the number is not at its correct position, swap
        if nums[i] != nums[correctPos] {
            nums[i], nums[correctPos] = nums[correctPos], nums[i]
        } else {
            // If the number is already at the correct position, move to the next element
            i++
        }
    }
}

// Find the missing number in an array containing n distinct numbers in the range [0, n]
func findMissingNumber(nums []int) int {
    i, n := 0, len(nums)
    
    // Cyclic sort (note that array contains 0 to n, with one number missing)
    for i < n {
        // If the current number is less than n and not at its correct position
        if nums[i] < n && nums[i] != i {
            // Swap with the correct position
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            // If the number is already at the correct position or is equal to n, move to the next element
            i++
        }
    }
    
    // Find the first missing number
    for i := 0; i < n; i++ {
        if nums[i] != i {
            return i
        }
    }
    
    // If all numbers from 0 to n-1 are present, the missing number is n
    return n
}

// Find all missing numbers in an array containing numbers from 1 to n
func findAllMissingNumbers(nums []int) []int {
    i := 0
    for i < len(nums) {
        // Correct position for the current number
        correctPos := nums[i] - 1
        
        // If the number is not at its correct position, swap
        if nums[i] != nums[correctPos] {
            nums[i], nums[correctPos] = nums[correctPos], nums[i]
        } else {
            // If the number is already at the correct position, move to the next element
            i++
        }
    }
    
    missingNumbers := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] != i + 1 {
            missingNumbers = append(missingNumbers, i + 1)
        }
    }
    
    return missingNumbers
}
```

## Example Problems

1. **Cyclic Sort**
2. **Find the Missing Number**
3. **Find All Missing Numbers**
4. **Find the Duplicate Number**
5. **Find All Duplicates**
6. **Find the Corrupt Pair**
7. **Find the Smallest Missing Positive Number**
8. **Find the First K Missing Positive Numbers**

Each of these problems has a dedicated solution file in this directory. 