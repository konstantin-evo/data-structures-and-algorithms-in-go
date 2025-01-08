<p align="center">
  <img src="src/img/header.png" alt="Logo" width="70%">
</p>

This is an educational repository based on the course "Data Structures & Algorithms In Go". The repository includes
practical implementations of various data structures and algorithms using the Go programming language. Each section
contains detailed guides, hands-on coding exercises, and thorough explanations of the principles at hand.

## Table of Contents

1. [Arrays](#arrays)
2. [Recursive Functions](#recursive-functions)
3. [Stack](#stack-operations)
4. Queue
5. [Tree](#tree-operations)
6. Binary Search Tree
7. [Sorting](#sorting)

## Arrays

| No.  | Task                                  | Description                                                                                                                        | Complexity | How to Run                                                               |
|------|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|------------|--------------------------------------------------------------------------|
| 1-01 | Sum Array                             | Sums all elements in an array.                                                                                                     | O(n)       | `cd 01-arrays/01-sum-array` <br> `go run main.go`                        |
| 1-02 | Sequential Search                     | Performs a linear search on an array.                                                                                              | O(n)       | `cd 01-arrays/02-sequential search` <br> `go run main.go`                |
| 1-03 | Binary Search                         | Performs a binary search on a sorted array.                                                                                        | O(log n)   | `cd 01-arrays/03-binary-search` <br> `go run main.go`                    |
| 1-04 | Max Subarray Sum (Kadane's Algorithm) | Finds the maximum sum of a subarray.                                                                                               | O(n)       | `cd 01-arrays/04-max-subarray-sum` <br> `go run main.go`                 |
| 1-05 | Array Rotation (Reversal method)      | Rotates an array by k positions.                                                                                                   | O(n)       | `cd 01-arrays/05-rotate-an-array` <br> `go run main.go`                  |
| 1-06 | Wave Array                            | Converts an array into a wave-like pattern.                                                                                        | O(n)       | `cd 01-arrays/06-wave-array` <br> `go run main.go`                       |
| 1-07 | Index Array                           | Reorders an array such that the value at each index matches its corresponding index.                                               | O(n)       | `cd 01-arrays/07-index-array` <br> `go run main.go`                      |
| 1-08 | Sorting Array (Cycle method)          | Sorts an array containing numbers from 1 to n in ascending order.                                                                  | O(n²)      | `cd 01-arrays/08-sorting-from-1-to-n` <br> `go run main.go`              |
| 1-09 | Smallest Positive Missing Number      | Finds the smallest positive missing number in an array.                                                                            | O(n)       | `cd 01-arrays/09-smallest-positive-missing-number` <br> `go run main.go` |
| 1-10 | Maximum Minimum Array                 | Reorders an array such that maximum and minimum elements alternate.                                                                | O(n)       | `cd 01-arrays/10-maximum-minimum-array` <br> `go run main.go`            |
| 1-11 | Array Index Maximum Difference        | Finds the maximum difference between indices where the element at the higher index is greater than the element at the lower index. | O(n)       | `cd 01-arrays/11-array-index-maximum-difference` <br> `go run main.go`   |

## Recursive Functions

| No.  | Task                           | Description                                                              | Complexity       | How to Run                                                                     |
|------|--------------------------------|--------------------------------------------------------------------------|------------------|--------------------------------------------------------------------------------|
| 2-01 | Convert Decimal to Hexadecimal | Recursively converts and prints a decimal integer in hexadecimal format. | O(log n)         | `cd 02-recursion/01-print-base-16-integers` <br> `go run main.go`              |
| 2-02 | GCD Calculation                | Recursively calculates the greatest common divisor of two integers.      | O(log min(m, n)) | `cd 02-recursion/02-greatest-common-divisor` <br> `go run main.go`             |
| 2-03 | Fibonacci Calculation          | Recursively calculates the nth Fibonacci number.                         | O(2^n)           | `cd 02-recursion/03-fibonacci-numbers` <br> `go run main.go`                   |
| 2-04 | All Permutations of a List     | Recursively generates and prints all permutations of an integer list.    | O(n!)            | `cd 02-recursion/04-all-permutations-of-an-integer-list` <br> `go run main.go` |
| 2-05 | Tower of Hanoi                 | Recursively solves the Tower of Hanoi problem for n disks.               | O(2^n)           | `cd 02-recursion/05-tower-of-hanoi` <br> `go run main.go`                      |
| 2-06 | Binary Search                  | Recursively searches for a value in a sorted array.                      | O(log n)         | `cd 02-recursion/06-binary-search` <br> `go run main.go`                       |

## Stack Operations

| No.  | Task                    | Description                                                      | Complexity                                                                                                      | How to Run                                                           |
|------|-------------------------|------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------|
| 3-01 | Stack Using Array       | Implements a stack using Array                                   | Push(): O(1) <br> Pop(): O(1) <br> Top(): O(1) <br> IsEmpty(): O(1) <br> Length(): O(1) <br> Print(): O(n) <br> | `cd 03-stack/01-stack-using-array` <br> `go run main.go`             |
| 3-02 | Stack Using Linked List | Implements a stack using Linked List                             | Same as above                                                                                                   | `cd 03-stack/02-stack-using-linked-list` <br> `go run main.go`       |
| 3-03 | Sorted Insert           | Inserts an element in sorted order                               | O(n)                                                                                                            | `cd 03-stack/03-sorted-insert` <br> `go run main.go`                 |
| 3-04 | Sort Stack              | Sorts a stack in ascending order                                 | O(n^2)                                                                                                          | `cd 03-stack/04-sort-stack` <br> `go run main.go`                    |
| 3-05 | Bottom Insert           | Inserts an element at the bottom of the stack                    | O(n)                                                                                                            | `cd 03-stack/05-bottom-insert` <br> `go run main.go`                 |
| 3-06 | Reverse Stack           | Reverses the order of elements in a stack                        | O(n²)                                                                                                           | `cd 03-stack/06-reverse-stack` <br> `go run main.go`                 |
| 3-07 | Reverse K Elements      | Reverses the first k elements in a stack                         | O(n)                                                                                                            | `cd 03-stack/07-reverse-k-elements-in-a-stack` <br> `go run main.go` |
| 3-08 | Balanced Parentheses    | Checks if parentheses in an expression are balanced <sup>*</sup> | O(n)                                                                                                            | `cd 03-stack/08-balanced-parentheses` <br> `go run main.go`          |

<sup>*</sup> Tasks with a note contain additional explanations.

---

### Balanced Parentheses

The closing symbol should be matched with the most recently seen opening symbol. For example, `{()}` is
legal, `{() ({})}` is legal, but `{((}` and `{(})` are not legal.

**Input**  
A string consisting solely of `(`, `)`, `{`, `}`, `[`, and `]`.

**Output**  
Return `true` if the parentheses are balanced. Otherwise, return `false`.

## Tree Operations

| No.  | Task                                      | Description                                                        | Complexity | How to Run                                                                                   |
|------|-------------------------------------------|--------------------------------------------------------------------|------------|----------------------------------------------------------------------------------------------|
| 5-01 | Complete Binary Tree                      | Constructs a complete binary tree from an array using level-order  | O(n)       | `cd 05-tree/01-complete-binary-tree` <br> `go run main.go`                                   |
| 5-02 | PreOrder Traversal                        | Implements and prints the preorder traversal of a binary tree      | O(n)       | `cd 05-tree/02-print-pre-order-traversal` <br> `go run main.go`                              |
| 5-03 | PostOrder Traversal                       | Implements and prints the postorder traversal of a binary tree     | O(n)       | `cd 05-tree/03-print-post-order-traversal` <br> `go run main.go`                             |
| 5-04 | InOrder Traversal                         | Implements and prints the inorder traversal of a binary tree       | O(n)       | `cd 05-tree/04-print-in-order-traversal` <br> `go run main.go`                               |
| 5-05 | Breadth-First Traversal                   | Implements and prints the level-order traversal of a binary tree   | O(n)       | `cd 05-tree/05-print-level-order-traversal` <br> `go run main.go`                            |
| 5-06 | DepthFirst Traversal                      | Implements and prints the depth-first traversal without recursion  | O(n)       | `cd 05-tree/06-print-depth-first-without-recursion` <br> `go run main.go`                    |
| 5-07 | Level-Order Line-by-Line Using Two Queues | Implements level-order traversal line by line using two queues     | O(n)       | `cd 05-tree/07-print-level-order-line-by-line-using-two-queues` <br> `go run main.go`        |
| 5-08 | Level-Order Line-by-Line Using One Queue  | Implements level-order traversal line by line using a single queue | O(n)       | `cd 05-tree/08-print-a-tree-in-level-order-line-by-line-using-a-queue` <br> `go run main.go` |

### Comparison of tree traversals

Example Tree Structure:

```
          1
        /   \
       2     3
      / \   / \
     4   5 6   7
```

| **Traversal Type**                        | **Description**                                                                             | **Order**                   | **Example Output** (Given Tree: 1 → 2, 3 → 4, 5, 6, 7) | **Type** |
|-------------------------------------------|---------------------------------------------------------------------------------------------|-----------------------------|--------------------------------------------------------|----------|
| **PreOrder Traversal**                    | Visit the root, then recursively traverse the left subtree, followed by the right subtree.  | Root → Left → Right         | **1, 2, 4, 5, 3, 6, 7**                                | **DFS**  |
| **InOrder Traversal**                     | Recursively traverse the left subtree, visit the root, then the right subtree.              | Left → Root → Right         | **4, 2, 5, 1, 6, 3, 7**                                | **DFS**  |
| **PostOrder Traversal**                   | Recursively traverse the left subtree, then the right subtree, and visit the root.          | Left → Right → Root         | **4, 5, 2, 6, 7, 3, 1**                                | **DFS**  |
| **Breadth-First Traversal (Level Order)** | Traverse the tree level by level, starting at the root and moving to each subsequent level. | Level 1 → Level 2 → Level 3 | **1, 2, 3, 4, 5, 6, 7**                                | **BFS**  |

This makes it clear which traversal types are part of **DFS** or **BFS**.
Observations:

- **PreOrder** is useful for creating a copy of the tree.
- **InOrder** is commonly used for binary search trees to retrieve elements in sorted order.
- **PostOrder** is helpful for deleting or freeing all nodes in a tree.
- **Breadth-First (Level Order)** provides an overview of the tree level by level.

## Sorting

| No.  | Task                 | Description                                                     | Complexity | How to Run                                                    |
|------|----------------------|-----------------------------------------------------------------|------------|---------------------------------------------------------------|
| 7-01 | Bubble Sort          | Implements bubble sort with ascending and descending options    | O(n^2)     | `cd 07-sorting/01-bubble-sort` <br> `go run main.go`          |
| 7-02 | Modified Bubble Sort | Implements modified bubble sort with optimized performance      | O(n^2)     | `cd 07-sorting/02-modified-bubble-sort` <br> `go run main.go` |
| 7-03 | Insertion Sort       | Implements insertion sort with ascending and descending options | O(n^2)     | `cd 07-sorting/03-insertion-sort` <br> `go run main.go`       |
| 7-04 | Selection Sort       | Implements selection sort with ascending option                 | O(n^2)     | `cd 07-sorting/04-selection-sort` <br> `go run main.go`       |
