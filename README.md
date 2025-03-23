# DSA Patterns Roadmap

A comprehensive Data Structures and Algorithms (DSA) learning resource with pattern-based approach. This repository contains implementations of common DSA patterns in Golang with detailed explanations, example problems, and solutions.

## Overview

This repository is organized around the concept of patterns in DSA. Each pattern represents a common approach or technique used to solve a specific class of problems. By mastering these patterns, you'll be able to recognize and solve a wide variety of algorithmic challenges.

Think of these patterns as tools in your toolbox - each designed for a specific type of problem. Learning when to use each tool is just as important as knowing how it works!

## Patterns Covered

1. **Sliding Window**
   - **Core Concept**: Maintain a "window" that slides through an array/string to process data in chunks.
   - **Simple Explanation**: Imagine looking through a moving window frame that shows only part of a long line of items. As you slide the window, you see new items entering from one side and old items leaving from the other side.
   - **Real-world Analogy**: Like watching scenery through a train window - you see new scenery come into view while old scenery disappears.
   - **When to Use**: 
     - Contiguous subarrays or substrings with constraints on size or properties
     - Finding maximum/minimum sum subarray of fixed size
     - Finding longest/shortest substring with specific properties
   - **Common Problems**: 
     - Maximum sum subarray of size K
     - Longest substring with K distinct characters
     - Fruits into baskets
     - Longest substring without repeating characters
   - **Visualization**: `[a b c d e f g]` → First window `[a b c]`, then `[b c d]`, then `[c d e]`, etc.
   - **Time Complexity**: Typically O(N) where N is the size of the array/string.
   
2. **Two Pointers**
   - **Core Concept**: Use two pointers to traverse data structures, usually moving toward or away from each other.
   - **Simple Explanation**: Place two fingers at different positions in a list and move them based on certain rules until you find the answer.
   - **Real-world Analogy**: Like closing in on a target from two sides or searching a room with two people starting from opposite ends.
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
   - **Visualization**: `[1 2 3 4 5]` → Pointers at `[1 ... 5]`, then `[2 ... 4]`, then `[3 ... 3]`
   - **Time Complexity**: Usually O(N) compared to O(N²) brute force approaches.
   
3. **Fast & Slow Pointers**
   - **Core Concept**: Use two pointers that move at different speeds through a sequence.
   - **Simple Explanation**: Like a race between a tortoise and a hare - one moves one step at a time, the other moves two steps.
   - **Real-world Analogy**: Similar to how a fast runner will eventually catch up to a slower runner on a circular track.
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
   - **Visualization**: `1→2→3→4→5→6→3(cycle)` → Slow pointer: `1→2→3`, Fast pointer: `1→3→5→4→6→3` (they meet at 3)
   - **Time Complexity**: O(N) where N is the length of the sequence.
   
4. **Merge Intervals**
   - **Core Concept**: Deal with overlapping intervals and merging them when needed.
   - **Simple Explanation**: Combine time slots that overlap to create a simpler schedule.
   - **Real-world Analogy**: Like merging multiple calendar appointments that overlap into single blocks of busy time.
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
   - **Visualization**: `[1,3], [2,6], [8,10]` → Merge to `[1,6], [8,10]`
   - **Time Complexity**: O(N log N) due to sorting of intervals.
   
5. **Cyclic Sort**
   - **Core Concept**: Sort arrays containing values in a given range by placing each element at its correct index.
   - **Simple Explanation**: If numbers are from 1 to N, place each number exactly where it belongs (at index number-1).
   - **Real-world Analogy**: Like organizing books on a shelf where each book has its specific spot based on its number.
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
   - **Visualization**: `[3,1,5,4,2]` → `[1,2,3,4,5]` by swapping elements to their correct positions
   - **Time Complexity**: O(N) as each element is moved at most once.
   
6. **In-place Reversal of a LinkedList**
   - **Core Concept**: Reverse a linked list or parts of it by changing the pointers in-place.
   - **Simple Explanation**: Flip the direction of the arrows connecting the nodes.
   - **Real-world Analogy**: Like reversing the direction of a one-way street by changing all the signs.
   - **When to Use**: 
     - Reverse entire linked list or sublist
     - Rotate linked list
     - Problems involving changing list structure
   - **Common Problems**: 
     - Reverse a linked list
     - Reverse a sublist
     - Reverse every K-element sublist
     - Rotate a linked list by K places
   - **Visualization**: `1→2→3→4→5` becomes `5→4→3→2→1` by reversing pointers
   - **Time Complexity**: O(N) where N is the number of nodes.
   
7. **Tree Breadth First Search (BFS)**
   - **Core Concept**: Traverse a tree level-by-level using a queue.
   - **Simple Explanation**: Visit all nodes at one level before moving to the next level.
   - **Real-world Analogy**: Like exploring a building floor by floor, completely covering one floor before moving to the next.
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
   - **Visualization**: For tree `(1(2,3(4,5)))`, BFS order is `1, 2, 3, 4, 5`
   - **Time Complexity**: O(N) where N is the number of nodes.
   
8. **Tree Depth First Search (DFS)**
   - **Core Concept**: Explore as far as possible along branches before backtracking.
   - **Simple Explanation**: Dive deep into one path of the tree before exploring other paths.
   - **Real-world Analogy**: Like exploring a maze by following one path until dead end before backtracking to try other paths.
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
   - **Visualization**: For tree `(1(2,3(4,5)))`, DFS preorder is `1, 2, 3, 4, 5`
   - **Time Complexity**: O(N) where N is the number of nodes.
   
9. **Two Heaps**
   - **Core Concept**: Use two heaps (a min-heap and a max-heap) to efficiently track medians or a partition point.
   - **Simple Explanation**: Divide your data into two groups - smaller half and larger half - to quickly access the middle or extremes.
   - **Real-world Analogy**: Like a balance scale with two sides that you maintain in equilibrium.
   - **When to Use**: 
     - Finding median from a data stream
     - Finding the balance point in a data set
     - Problems requiring partitioning elements around a specific point
   - **Common Problems**: 
     - Find the median of a number stream
     - Sliding window median
     - IPO (maximize capital)
     - Next interval
   - **Visualization**: `max-heap [1,2,3] | min-heap [4,5,6]` with median being 3 or 4
   - **Time Complexity**: O(log N) per element insertion/deletion.
   
10. **Subsets**
    - **Core Concept**: Generate all possible subsets, permutations, or combinations of a set.
    - **Simple Explanation**: Create all possible ways to select or arrange items from a group.
    - **Real-world Analogy**: Like figuring out all possible outfits from a set of clothes, or all possible pizza toppings from available options.
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
    - **Visualization**: For set `[1,2]`, subsets are `[], [1], [2], [1,2]`
    - **Time Complexity**: Typically O(2^N) for subsets, O(N!) for permutations.
    
11. **Modified Binary Search**
    - **Core Concept**: Adapt binary search for more complex scenarios beyond standard search.
    - **Simple Explanation**: Use divide-and-conquer to efficiently search in sorted arrays, even with special conditions.
    - **Real-world Analogy**: Like searching for a name in a phone book but with tricks like missing pages or multiple copies.
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
    - **Visualization**: `[1,2,3,4,5,6,7]` → check middle (4), then search left or right half based on comparison
    - **Time Complexity**: O(log N) where N is the size of the array.
    
12. **Bitwise XOR**
    - **Core Concept**: Use XOR operations to solve problems efficiently.
    - **Simple Explanation**: XOR is like a toggle switch - it flips a bit if the corresponding bit is 1, unchanged if 0.
    - **Real-world Analogy**: Like a light switch that toggles state each time it's pressed - press twice and it returns to original state.
    - **When to Use**: 
      - Finding single numbers among duplicates
      - Missing or duplicate numbers
      - Problems involving bit manipulation
    - **Common Problems**: 
      - Single number
      - Two single numbers
      - Complement of base 10 number
      - Flip image
    - **Visualization**: `5 XOR 5 = 0` (since 101 XOR 101 = 000)
    - **Time Complexity**: Typically O(N) with constant space.
    
13. **Top K Elements**
    - **Core Concept**: Find or maintain the K largest/smallest elements using a heap data structure.
    - **Simple Explanation**: Keep track of the "top K" items in a special container that automatically keeps them in order.
    - **Real-world Analogy**: Like maintaining a leaderboard of the top K scores in a game.
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
    - **Visualization**: For array `[3,1,5,12,2,11]` and K=3, top K largest are `[5,11,12]`
    - **Time Complexity**: Typically O(N log K) where N is input size and K is the parameter.
    
14. **K-way Merge**
    - **Core Concept**: Merge K sorted arrays/lists efficiently using a min-heap.
    - **Simple Explanation**: Combine multiple sorted lists into one sorted list by picking the smallest item each time.
    - **Real-world Analogy**: Like merging multiple sorted piles of cards into one sorted pile.
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
    - **Visualization**: Merging `[1,4,7], [2,5,8], [3,6,9]` into `[1,2,3,4,5,6,7,8,9]`
    - **Time Complexity**: O(N log K) where N is total number of elements and K is the number of lists.
    
15. **Dynamic Programming (0/1 Knapsack)**
    - **Core Concept**: Solve optimization problems by making selections from a set of items.
    - **Simple Explanation**: Choose the best items to maximize value while staying within constraints.
    - **Real-world Analogy**: Like packing a backpack for a hike - you need to decide which items to take to maximize usefulness while not exceeding weight limit.
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
    - **Visualization**: For items `(value,weight)` as `(1,1), (4,3), (5,4), (7,5)` with capacity 7, optimal selection is `(4,3)` and `(5,4)`
    - **Time Complexity**: Typically O(N*W) where N is the number of items and W is the capacity.
    
16. **Topological Sort**
    - **Core Concept**: Order vertices in a directed acyclic graph (DAG) such that all directed edges go from earlier to later vertices.
    - **Simple Explanation**: Arrange tasks in order so that all prerequisites are completed before dependent tasks.
    - **Real-world Analogy**: Like planning the order of university courses where some courses require prerequisites.
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
    - **Visualization**: For dependencies A→B, A→C, B→D, C→D, order is `A,B,C,D` or `A,C,B,D`
    - **Time Complexity**: O(V+E) where V is the number of vertices and E is the number of edges.

## Additional Patterns Worth Learning

17. **Dynamic Programming (Fibonacci Numbers)**
    - **Core Concept**: Use previous computations to solve problems with overlapping subproblems.
    - **Simple Explanation**: Break down problems into smaller parts and remember the solutions to avoid recalculating.
    - **Real-world Analogy**: Like writing down math steps to avoid redoing calculations you've already done.
    - **When to Use**: Problems with Fibonacci-like recurrence relations.
    - **Common Problems**: Fibonacci sequence, climbing stairs, house thief.
    - **Visualization**: Calculating Fibonacci: F(4) = F(3) + F(2) = (F(2) + F(1)) + F(2) = ...

18. **Union Find (Disjoint Set)**
    - **Core Concept**: Track connected components in undirected graphs.
    - **Simple Explanation**: Group items together and quickly check if two items belong to the same group.
    - **Real-world Analogy**: Like tracking friend circles - if A is friends with B and B is friends with C, then A is in the same friend circle as C.
    - **When to Use**: Network connectivity, redundant connections, minimum spanning tree.
    - **Common Problems**: Friend circles, number of islands, redundant connection.
    - **Visualization**: Groups: `{1,2,5}, {3,6}, {4}` → Union(2,6) → `{1,2,5,3,6}, {4}`

19. **Greedy Algorithms**
    - **Core Concept**: Make locally optimal choices at each stage.
    - **Simple Explanation**: Always pick the currently best option without looking ahead.
    - **Real-world Analogy**: Like always picking the shortest checkout line at a store without considering how many items each person has.
    - **When to Use**: Problems where local optimality leads to global optimality.
    - **Common Problems**: Activity selection, coin change, interval scheduling.
    - **Visualization**: For coin change with coins [25,10,5,1] and amount 36, greedy takes 25+10+1 = 36

20. **Tries**
    - **Core Concept**: Tree-like data structure for storing strings.
    - **Simple Explanation**: A word tree where each path from root to a node spells out a word or prefix.
    - **Real-world Analogy**: Like a dictionary where words with the same prefix share the same starting path.
    - **When to Use**: Word search, prefix matching, autocompletion.
    - **Common Problems**: Implement Trie, word search, longest common prefix.
    - **Visualization**: Trie for "cat", "car", "dog": root→c→a→t/r, root→d→o→g

## Common Mistakes to Avoid

- **Sliding Window**: Forgetting to update window bounds correctly, or using it on non-contiguous problems
- **Two Pointers**: Incorrect pointer movement rules, or applying to unsorted arrays when sorting is required
- **Fast & Slow Pointers**: Not handling the null pointer cases, or incorrect cycle detection logic
- **Merge Intervals**: Forgetting to sort intervals first, or incorrect merging conditions
- **Cyclic Sort**: Using with numbers outside the range [1...n], or incorrect indexing
- **Tree Traversal**: Confusing BFS and DFS, or improper queue/stack management
- **Dynamic Programming**: Incorrect state definition, or missing base cases
- **Topological Sort**: Applying to graphs with cycles, or incorrect dependency representation

## Problem-Solving Framework

1. **Understand**: Read the problem carefully and identify the key constraints
2. **Match**: Determine which pattern(s) might be applicable
3. **Approach**: Design a solution using the appropriate pattern(s)
4. **Code**: Implement the solution using the pattern template
5. **Test**: Verify with examples and edge cases
6. **Optimize**: Look for opportunities to improve time/space complexity

## How to Use This Repository

Each pattern has its own directory containing:
1. A README.md with a detailed explanation of the pattern, when to use it, and its variations
2. Golang code implementations of the core algorithms
3. Example problems that can be solved using the pattern
4. Step-by-step solutions with time and space complexity analysis
5. Test cases to verify the implementations
6. Visual explanations and diagrams where applicable

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