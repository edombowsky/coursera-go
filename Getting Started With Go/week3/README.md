# Module 3 Activity

[Assignment](https://www.coursera.org/learn/golang-getting-started/peer/sLPZg/module-3-activity-slice-go)

## Instructions

The goal of this assignment is to practice working with integers, slices and loops in Go.

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

## My assignment

Source code at `slice.go`

## Testing

```text
Enter an integer, X to quit: 1
After insert: [1]
Enter an integer, X to quit: 2
After insert: [1 2]
Enter an integer, X to quit: 4
After insert: [1 2 4]
Enter an integer, X to quit: 5
After insert: [1 2 4 5]
Enter an integer, X to quit: 7
After insert: [1 2 4 5 7]
Enter an integer, X to quit: -1
After insert: [-1 1 2 4 5 7]
Enter an integer, X to quit: 3
After insert: [-1 1 2 3 4 5 7]
Enter an integer, X to quit: -10
After insert: [-10 -1 1 2 3 4 5 7]
Enter an integer, X to quit: x
```
