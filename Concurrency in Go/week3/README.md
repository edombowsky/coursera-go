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

Enter a sequence of integers (more than 4), X to quit: 999 1000 1 2 3 4 -9 -345 100 102 7 78 0 54

--- Unsorted --- 

 [999 1000 1 2 3 4 -9 -345 100 102 7 78 0 54]

Total number of integers entered [14]
Number of integers in each slice [3]

Initial slice of integers to be sorted in Thread #1: [999 1000 1 2]
Initial slice of integers to be sorted in Thread #2: [3 4 -9 -345]
Initial slice of integers to be sorted in Thread #3: [100 102 7]
Initial slice of integers to be sorted in Thread #4: [78 0 54]
On thread #[4] - Slice received [78 0 54]
On thread #[4] - Sorting ...
On thread #[4] - Slice sorted [0 54 78]. Exiting thread...
On thread #[1] - Slice received [999 1000 1 2]
On thread #[1] - Sorting ...
On thread #[1] - Slice sorted [1 2 999 1000]. Exiting thread...
On thread #[2] - Slice received [3 4 -9 -345]
On thread #[2] - Sorting ...
On thread #[2] - Slice sorted [-345 -9 3 4]. Exiting thread...
On thread #[3] - Slice received [100 102 7]
On thread #[3] - Sorting ...
On thread #[3] - Slice sorted [7 100 102]. Exiting thread...

Resulting slice of integers sorted in Thread #1: [1 2 999 1000]
Resulting slice of integers sorted in Thread #2: [-345 -9 3 4]
Resulting slice of integers sorted in Thread #3: [7 100 102]
Resulting slice of integers sorted in Thread #4: [0 54 78]

After merging all that slices, the initial slice of integers now contains...: [1 2 999 1000 -345 -9 3 4 7 100 102 0 54 78]

--- Sorted ---

 [-345 -9 0 1 2 3 4 7 54 78 100 102 999 1000]
```
