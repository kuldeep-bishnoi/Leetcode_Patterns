# Bitwise XOR Pattern

## Introduction

The Bitwise XOR pattern utilizes the XOR (exclusive OR) operation to solve various problems efficiently. XOR is a binary operation that takes two equal-length bit patterns and performs the logical XOR operation on each pair of corresponding bits. The result in each position is 1 if only one of the bits is 1, but 0 if both are 0 or both are 1.

In Golang, the XOR operation is represented by the `^` operator.

## How It Works

XOR has several mathematical properties that make it useful in algorithm design:

1. **Commutative**: `a ^ b = b ^ a`
2. **Associative**: `a ^ (b ^ c) = (a ^ b) ^ c`
3. **Identity element**: `a ^ 0 = a`
4. **Self-inverse**: `a ^ a = 0`

The self-inverse property is particularly useful. If we XOR a number with itself, we get 0. This property can be leveraged to find a missing number or a single non-duplicate number in a sequence.

## Time and Space Complexity

- **Time Complexity**: Most solutions using the bitwise XOR pattern have O(n) time complexity where n is the size of the input array.
- **Space Complexity**: Typically O(1), as XOR operations don't require additional space proportional to the input size.

## When to Use Bitwise XOR Pattern

This pattern is particularly useful when:

1. You need to find a missing number in a sequence
2. You need to find a single non-duplicate number among duplicates
3. You need to swap two numbers without using a temporary variable
4. You need to detect if exactly one bit is set
5. You need to find all pairs with a specific XOR value
6. You need to execute operations like finding the complement of a number

## Common Problem Patterns

1. **Find the missing number**: Given an array of n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing.
2. **Find the single number**: Given a non-empty array where every element appears twice except for one, find that single element.
3. **Find two missing numbers**: Given an array of n-2 distinct numbers taken from 0, 1, 2, ..., n, find the two that are missing.
4. **Find two single numbers**: Given an array where all numbers appear twice except two numbers, find those two numbers.
5. **Complement of a number**: Find the complement of a given number (flipping all bits).
6. **Flip and invert an image**: Given a binary matrix representing an image, flip it horizontally, then invert it.

## Implementation in Golang

Here's a simple example of how to use XOR to find a single non-duplicate number in an array where all other numbers appear twice:

```go
func findSingleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

## Example Problems

1. **Find the Missing Number**: Find the missing number in an array containing n distinct numbers taken from 0 to n.
2. **Single Number**: Find the element that appears once in an array where every other element appears twice.
3. **Single Number III**: Find two elements that appear only once in an array where all other elements appear exactly twice.
4. **Complement of Base 10 Number**: Find the complement of a given decimal number.
5. **Flipping an Image**: Flip and invert a binary matrix.

Each problem demonstrates the power and efficiency of using bitwise XOR operations to solve what would otherwise be more complex problems. 