# Module 3 Activity - THREADS IN GO

[Assignment](https://www.coursera.org/learn/golang-concurrency/peer/meeiu/module-3-activity)

## Instructions

The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four `goroutines` to create four sub-arrays and then merge the arrays into a single array.

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts 1/4 of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

## My assignment

Source code at `sort.go`

## Testing

```text
go run sort.go
```
