# Modified Binary Search Pattern

## Introduction

The Modified Binary Search pattern is an extension of the traditional binary search algorithm, adapted to solve problems where a sorted array is given but the search condition is more complex than a simple equality check. This pattern is particularly useful when dealing with rotated, shifted, or otherwise modified sorted arrays, or when searching for elements that meet specific criteria rather than exact matches.

## How It Works

The standard binary search algorithm works by repeatedly dividing the search interval in half to efficiently find an element in a sorted array. The Modified Binary Search pattern builds on this foundation but introduces variations to handle more complex scenarios:

1. **Identify the Sorted Half**: In cases with rotated arrays, determine which half of the array is sorted.
2. **Check Special Conditions**: Instead of just checking for equality, examine if a mid-point meets certain criteria.
3. **Adjust Search Space**: Modify how the search space is adjusted based on the specific problem constraints.
4. **Handle Duplicates**: Incorporate special handling for arrays with duplicate elements when necessary.

## Time and Space Complexity

- **Time Complexity**: O(log n) in most cases, where n is the size of the input array. This logarithmic time complexity is what makes binary search and its variants so efficient.
- **Space Complexity**: O(1) for iterative implementations, or O(log n) for recursive implementations due to the function call stack.

## When to Use Modified Binary Search Pattern

This pattern is appropriate when:

- You're working with a sorted array (or partially sorted array like a rotated sorted array).
- The problem involves finding a specific element or a range that meets certain criteria.
- You need an efficient solution with logarithmic time complexity.
- The problem statement hints at "search," "find," or "locate" in a sorted structure.

## Common Problem Patterns

1. **Search in Rotated Sorted Array**: Finding an element in an array that has been rotated.
2. **Search in a Sorted Array of Unknown Size**: Efficiently searching when array bounds are not known.
3. **Ceiling of a Number**: Finding the smallest element in the array greater than or equal to a target.
4. **Floor of a Number**: Finding the largest element in the array less than or equal to a target.
5. **Next Letter**: Finding the smallest letter greater than a target.
6. **Number Range**: Finding the first and last position of an element.
7. **Search in a Sorted Infinite Array**: Efficiently search in an array of unknown size.
8. **Minimum Difference Element**: Finding the element with minimum difference with target.
9. **Bitonic Array Maximum**: Finding the maximum value in a bitonic array (increasing then decreasing).
10. **Peak Element**: Finding a peak element (larger than its neighbors).

## Implementation in Golang

A basic template for the Modified Binary Search pattern in Golang:

```go
func modifiedBinarySearch(arr []int, key int) int {
    start, end := 0, len(arr)-1
    
    for start <= end {
        mid := start + (end-start)/2
        
        if arr[mid] == key {
            return mid // Found the key
        }
        
        // Modified logic based on the specific problem
        if someCondition(arr, mid, key) {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    
    // Key not found or return value based on problem requirements
    return -1 // or another appropriate value
}
```

For rotated sorted arrays:

```go
func searchInRotatedArray(arr []int, key int) int {
    start, end := 0, len(arr)-1
    
    for start <= end {
        mid := start + (end-start)/2
        
        if arr[mid] == key {
            return mid
        }
        
        // Check which half is sorted
        if arr[start] <= arr[mid] { // Left half is sorted
            if key >= arr[start] && key < arr[mid] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        } else { // Right half is sorted
            if key > arr[mid] && key <= arr[end] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }
    
    return -1
}
```

## Example Problems

1. **Order-agnostic Binary Search**: Binary search when the sorting order is unknown.
2. **Search in a Rotated Sorted Array**: Find a target value in a rotated sorted array.
3. **Ceiling of a Number**: Find the smallest element greater than or equal to a target.
4. **Next Letter**: Find the smallest letter greater than a target letter.
5. **Number Range**: Find the first and last position of an element in a sorted array.
6. **Search in a Sorted Infinite Array**: Search in a sorted array of unknown size.
7. **Minimum Difference Element**: Find the element that has minimum difference with a target.
8. **Bitonic Array Maximum**: Find the maximum value in a bitonic array.
9. **Search Bitonic Array**: Search for a target in a bitonic array.
10. **Search in Row-wise and Column-wise Sorted Matrix**: Search in a 2D matrix sorted in both dimensions.

Each of these problems demonstrates a unique variation of the binary search algorithm, showing the versatility and power of the Modified Binary Search pattern. 