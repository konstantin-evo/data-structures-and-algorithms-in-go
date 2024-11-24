
<p align="center">
  <img src="src/img/header.png" alt="Logo" width="70%">
</p>

This is an educational repository based on the course "Data Structures & Algorithms In Go". The repository includes
practical implementations of various data structures and algorithms using the Go programming language. Each section
contains detailed guides, hands-on coding exercises, and thorough explanations of the principles at hand.

## Table of Contents

- [Arrays](#arrays)
- Recursive Functions
- Approaches to Solve Algorithm Design Problems
- Data Structures in Go
- Sorting
- Stack
- Queue
- Tree
- Binary Search Tree

## Arrays

| Task                                  | Description                                                                                                                        | Complexity | How to Run                                                               |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|------------|--------------------------------------------------------------------------|
| Sum Array                             | Sums all elements in an array.                                                                                                     | O(n)       | `cd 01-arrays/01-sum-array` <br> `go run main.go`                        |
| Sequential Search                     | Performs a linear search on an array.                                                                                              | O(n)       | `cd 01-arrays/02-sequential search` <br> `go run main.go`                |
| Binary Search                         | Performs a binary search on a sorted array.                                                                                        | O(log n)   | `cd 01-arrays/03-binary-search` <br> `go run main.go`                    |
| Max Subarray Sum (Kadane's Algorithm) | Finds the maximum sum of a subarray.                                                                                               | O(n)       | `cd 01-arrays/04-max-subarray-sum` <br> `go run main.go`                 |
| Array Rotation (Reversal method)      | Rotates an array by k positions.                                                                                                   | O(n)       | `cd 01-arrays/05-rotate-an-array` <br> `go run main.go`                  |
| Wave Array                            | Converts an array into a wave-like pattern.                                                                                        | O(n)       | `cd 01-arrays/06-wave-array` <br> `go run main.go`                       |
| Index Array                           | Reorders an array such that the value at each index matches its corresponding index.                                               | O(n)       | `cd 01-arrays/07-index-array` <br> `go run main.go`                      |
| Sorting Array (Cycle method)          | Sorts an array containing numbers from 1 to n in ascending order.                                                                  | O(n²)      | `cd 01-arrays/08-sorting-from-1-to-n` <br> `go run main.go`              |
| Smallest Positive Missing Number      | Finds the smallest positive missing number in an array.                                                                            | O(n)       | `cd 01-arrays/09-smallest-positive-missing-number` <br> `go run main.go` |
| Maximum Minimum Array                 | Reorders an array such that maximum and minimum elements alternate.                                                                | O(n)       | `cd 01-arrays/10-maximum-minimum-array` <br> `go run main.go`            |
| Array Index Maximum Difference        | Finds the maximum difference between indices where the element at the higher index is greater than the element at the lower index. | O(n)       | `cd 01-arrays/11-array-index-maximum-difference` <br> `go run main.go`   |

## Recursive Functions

| Task                           | Description                                                              | Complexity       | How to Run                                                                     |
|--------------------------------|--------------------------------------------------------------------------|------------------|--------------------------------------------------------------------------------|
| Convert Decimal to Hexadecimal | Recursively converts and prints a decimal integer in hexadecimal format. | O(log n)         | `cd 02-recursion/01-print-base-16-integers` <br> `go run main.go`              |
| GCD Calculation                | Recursively calculates the greatest common divisor of two integers.      | O(log min(m, n)) | `cd 02-recursion/02-greatest-common-divisor` <br> `go run main.go`             |
| Fibonacci Calculation          | Recursively calculates the nth Fibonacci number.                         | O(2^n)           | `cd 02-recursion/03-fibonacci-numbers` <br> `go run main.go`                   |
| All Permutations of a List     | Recursively generates and prints all permutations of an integer list.    | O(n!)            | `cd 02-recursion/04-all-permutations-of-an-integer-list` <br> `go run main.go` |
| Tower of Hanoi                 | Recursively solves the Tower of Hanoi problem for n disks.               | O(2^n)           | `cd 02-recursion/05-tower-of-hanoi` <br> `go run main.go`                      |
| Binary Search                  | Recursively searches for a value in a sorted array.                      | O(log n)         | `cd 02-recursion/06-binary-search` <br> `go run main.go`                       |
