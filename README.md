# DSA Patterns Roadmap

A comprehensive Data Structures and Algorithms (DSA) learning resource with pattern-based approach. This repository contains implementations of common DSA patterns in Golang with detailed explanations, example problems, and solutions.

## Overview

This repository is organized around the concept of patterns in DSA. Each pattern represents a common approach or technique used to solve a specific class of problems. By mastering these patterns, you'll be able to recognize and solve a wide variety of algorithmic challenges.

## Patterns Covered

1. **Sliding Window**
   - **Core Concept**: Maintain a "window" that slides through an array/string to process data in chunks.
   - **When to Use**: 
     - Contiguous subarrays or substrings with constraints on size or properties
     - Finding maximum/minimum sum subarray of fixed size
     - Finding longest/shortest substring with specific properties
   - **Common Problems**: 
     - Maximum sum subarray of size K
     - Longest substring with K distinct characters
     - Fruits into baskets
     - Longest substring without repeating characters
   - **Time Complexity**: Typically O(N) where N is the size of the array/string.
   
2. **Two Pointers**
   - **Core Concept**: Use two pointers to traverse data structures, usually moving toward or away from each other.
   - **When to Use**: 
     - Searching for pairs in a sorted array
     - Comparing values at different positions
     - Removing duplicates
     - Reversing arrays or strings
   - **Common Problems**: 
     - Pair with target sum
     - Remove duplicates
     - Squaring a sorted array
     - Triplet sum to zero (3Sum)
     - Dutch national flag problem (sort colors)
   - **Time Complexity**: Usually O(N) compared to O(N²) brute force approaches.
   
3. **Fast & Slow Pointers**
   - **Core Concept**: Use two pointers that move at different speeds through a sequence.
   - **When to Use**: 
     - Cycle detection in linked lists or arrays
     - Finding the middle of a linked list
     - Determining linked list length
     - Finding the start of a cycle
   - **Common Problems**: 
     - Linked list cycle detection
     - Start of linked list cycle
     - Happy number
     - Middle of the linked list
     - Palindrome linked list
   - **Time Complexity**: O(N) where N is the length of the sequence.
   
4. **Merge Intervals**
   - **Core Concept**: Deal with overlapping intervals and merging them when needed.
   - **When to Use**: 
     - Schedule/calendar problems
     - Interval merging/insertion
     - Finding free time slots
     - Resource allocation with time constraints
   - **Common Problems**: 
     - Merge intervals
     - Insert interval
     - Intervals intersection
     - Conflicting appointments
     - Minimum meeting rooms required
   - **Time Complexity**: O(N log N) due to sorting of intervals.
   
5. **Cyclic Sort**
   - **Core Concept**: Sort arrays containing values in a given range by placing each element at its correct index.
   - **When to Use**: 
     - Arrays containing numbers in range [1...n] with or without duplicates
     - Finding missing/duplicate numbers in a limited range
     - Problems where input array can be modified in-place
   - **Common Problems**: 
     - Missing number
     - Find all missing numbers
     - Find the duplicate number
     - Find all duplicates
     - Find the corrupt pair
   - **Time Complexity**: O(N) as each element is moved at most once.
   
6. **In-place Reversal of a LinkedList**
   - **Core Concept**: Reverse a linked list or parts of it by changing the pointers in-place.
   - **When to Use**: 
     - Reverse entire linked list or sublist
     - Rotate linked list
     - Problems involving changing list structure
   - **Common Problems**: 
     - Reverse a linked list
     - Reverse a sublist
     - Reverse every K-element sublist
     - Rotate a linked list by K places
   - **Time Complexity**: O(N) where N is the number of nodes.
   
7. **Tree Breadth First Search (BFS)**
   - **Core Concept**: Traverse a tree level-by-level using a queue.
   - **When to Use**: 
     - Level-order traversal
     - Finding shortest paths
     - Problems requiring level-wise processing
     - Connected components in graphs
   - **Common Problems**: 
     - Binary tree level order traversal
     - Zigzag traversal
     - Level averages
     - Minimum depth of a binary tree
     - Connect level order siblings
   - **Time Complexity**: O(N) where N is the number of nodes.
   
8. **Tree Depth First Search (DFS)**
   - **Core Concept**: Explore as far as possible along branches before backtracking.
   - **When to Use**: 
     - Path finding problems
     - Tree traversal (preorder, inorder, postorder)
     - Finding heights or depths
     - Problems requiring exploration of all paths
   - **Common Problems**: 
     - Path sum
     - All paths for a sum
     - Sum of path numbers
     - Path with maximum sum
     - Binary tree diameter
   - **Time Complexity**: O(N) where N is the number of nodes.
   
9. **Two Heaps**
   - **Core Concept**: Use two heaps (a min-heap and a max-heap) to efficiently track medians or a partition point.
   - **When to Use**: 
     - Finding median from a data stream
     - Finding the balance point in a data set
     - Problems requiring partitioning elements around a specific point
   - **Common Problems**: 
     - Find the median of a number stream
     - Sliding window median
     - IPO (maximize capital)
     - Next interval
   - **Time Complexity**: O(log N) per element insertion/deletion.
   
10. **Subsets**
    - **Core Concept**: Generate all possible subsets, permutations, or combinations of a set.
    - **When to Use**: 
      - Exhaustive search problems
      - Combinations and permutations
      - Problems requiring exploration of all possible options
    - **Common Problems**: 
      - Subsets with/without duplicates
      - Permutations with/without duplicates
      - String permutations by changing case
      - Generate parentheses
      - Unique generalized abbreviations
    - **Time Complexity**: Typically O(2^N) for subsets, O(N!) for permutations.
    
11. **Modified Binary Search**
    - **Core Concept**: Adapt binary search for more complex scenarios beyond standard search.
    - **When to Use**: 
      - Searching in sorted arrays with variations
      - Finding boundaries
      - Rotated or modified sorted arrays
      - Searching in sorted but infinite arrays
    - **Common Problems**: 
      - Order-agnostic binary search
      - Ceiling of a number
      - Next letter
      - Number range
      - Search in a sorted infinite array
      - Minimum difference element
      - Search in rotated array
    - **Time Complexity**: O(log N) where N is the size of the array.
    
12. **Bitwise XOR**
    - **Core Concept**: Use XOR operations to solve problems efficiently.
    - **When to Use**: 
      - Finding single numbers among duplicates
      - Missing or duplicate numbers
      - Problems involving bit manipulation
    - **Common Problems**: 
      - Single number
      - Two single numbers
      - Complement of base 10 number
      - Flip image
    - **Time Complexity**: Typically O(N) with constant space.
    
13. **Top K Elements**
    - **Core Concept**: Find or maintain the K largest/smallest elements using a heap data structure.
    - **When to Use**: 
      - Finding the K largest/smallest elements
      - Finding the K most frequent elements
      - Problems requiring partial sorting
      - When complete sorting is unnecessary
    - **Common Problems**: 
      - Kth smallest/largest number
      - K closest points to origin
      - Connect ropes with minimum cost
      - Top K frequent elements
      - Frequency sort
      - Kth closest numbers
    - **Time Complexity**: Typically O(N log K) where N is input size and K is the parameter.
    
14. **K-way Merge**
    - **Core Concept**: Merge K sorted arrays/lists efficiently using a min-heap.
    - **When to Use**: 
      - Merging K sorted arrays/lists
      - Problems involving multiple sorted sources
      - Finding smallest elements across multiple arrays
    - **Common Problems**: 
      - Merge K sorted lists
      - Kth smallest number in M sorted lists
      - Kth smallest element in a sorted matrix
      - Smallest number range
      - Find K pairs with smallest sums
    - **Time Complexity**: O(N log K) where N is total number of elements and K is the number of lists.
    
15. **Dynamic Programming (0/1 Knapsack)**
    - **Core Concept**: Solve optimization problems by making selections from a set of items.
    - **When to Use**: 
      - Binary decision problems (take/leave)
      - Problems with capacity constraints
      - Optimization problems with overlapping subproblems
    - **Common Problems**: 
      - 0/1 Knapsack
      - Subset sum
      - Equal subset sum partition
      - Minimum subset sum difference
      - Count of subset sum
      - Target sum
    - **Time Complexity**: Typically O(N*W) where N is the number of items and W is the capacity.
    
16. **Topological Sort**
    - **Core Concept**: Order vertices in a directed acyclic graph (DAG) such that all directed edges go from earlier to later vertices.
    - **When to Use**: 
      - Task scheduling with dependencies
      - Course prerequisites
      - Build systems
      - Any problem involving dependencies between entities
    - **Common Problems**: 
      - Task scheduling
      - Course schedule
      - Alien dictionary
      - Reconstructing a sequence
      - Minimum height trees
    - **Time Complexity**: O(V+E) where V is the number of vertices and E is the number of edges.

## Additional Patterns Worth Learning

17. **Dynamic Programming (Fibonacci Numbers)**
    - **Core Concept**: Use previous computations to solve problems with overlapping subproblems.
    - **When to Use**: Problems with Fibonacci-like recurrence relations.
    - **Common Problems**: Fibonacci sequence, climbing stairs, house thief.

18. **Union Find (Disjoint Set)**
    - **Core Concept**: Track connected components in undirected graphs.
    - **When to Use**: Network connectivity, redundant connections, minimum spanning tree.
    - **Common Problems**: Friend circles, number of islands, redundant connection.

19. **Greedy Algorithms**
    - **Core Concept**: Make locally optimal choices at each stage.
    - **When to Use**: Problems where local optimality leads to global optimality.
    - **Common Problems**: Activity selection, coin change, interval scheduling.

20. **Tries**
    - **Core Concept**: Tree-like data structure for storing strings.
    - **When to Use**: Word search, prefix matching, autocompletion.
    - **Common Problems**: Implement Trie, word search, longest common prefix.

## How to Use This Repository

Each pattern has its own directory containing:
1. A README.md with a detailed explanation of the pattern, when to use it, and its variations
2. Golang code implementations of the core algorithms
3. Example problems that can be solved using the pattern
4. Step-by-step solutions with time and space complexity analysis
5. Test cases to verify the implementations

## Learning Path

For beginners, I recommend following these patterns in order of increasing complexity:

1. **Fundamentals First**
   - Start with Two Pointers and Sliding Window
   - Practice Fast & Slow Pointers for linked list problems
   - Learn Merge Intervals and Cyclic Sort for array manipulations

2. **Tree Traversals**
   - Master Tree BFS for level-order traversals
   - Learn Tree DFS for path-related problems
   - Combine both for complex tree problems

3. **Intermediate Techniques**
   - Study Modified Binary Search variations
   - Implement Top K Elements with heaps
   - Learn K-way Merge for handling multiple sorted lists

4. **Advanced Patterns**
   - Tackle Topological Sort for dependency problems
   - Master Dynamic Programming patterns
   - Explore Bitwise XOR for optimization

5. **Expert Level**
   - Combine multiple patterns to solve complex problems
   - Optimize solutions for time and space efficiency
   - Create your own variations of these patterns

## Problem Recognition Tips

To identify which pattern to use:

- **Contiguous subarrays/substrings?** → Sliding Window
- **Sorted array or two data structures?** → Two Pointers
- **Linked list with cycle or finding middle?** → Fast & Slow Pointers
- **Time intervals or scheduling?** → Merge Intervals
- **Array with values in range [1...n]?** → Cyclic Sort
- **Tree traversal by level?** → BFS
- **Path or depth-related tree problems?** → DFS
- **Finding median or partition point?** → Two Heaps
- **All possible combinations/permutations?** → Subsets
- **Sorted array with modifications?** → Modified Binary Search
- **Top/smallest K elements?** → Heap
- **Multiple sorted arrays?** → K-way Merge
- **Optimization with constraints?** → Dynamic Programming
- **Dependencies or prerequisites?** → Topological Sort

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request to:
- Add new patterns or examples
- Improve existing implementations
- Enhance explanations or visualizations
- Fix bugs or optimize code

## License

This project is licensed under the MIT License - see the LICENSE file for details. 