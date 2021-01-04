# Module 1 Activity - CONCURRENCY BASICS

[Assignment](https://www.coursera.org/learn/golang-concurrency/peer/RAJ0V/module-2-activity)

## Instructions

The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.

Write two `goroutines` which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

## My assignment

Source code at `race.go`

## Testing

```text
## Example output 1:

go run race.go
dec:  4 index:  0
inc:  5 index:  0
dec:  4 index:  1
inc:  4 index:  1
dec:  3 index:  2
dec:  3 index:  3
inc:  4 index:  2
dec:  3 index:  4
dec:  2 index:  5
inc:  3 index:  3
dec:  2 index:  6
dec:  1 index:  7

// Example output 2:

go run race.go
inc:  6 index:  0
dec:  5 index:  0
dec:  4 index:  1
inc:  5 index:  1
dec:  4 index:  2
dec:  3 index:  3
dec:  2 index:  4
inc:  3 index:  2
dec:  2 index:  5
dec:  1 index:  6
inc:  2 index:  3
dec:  1 index:  7
```

## Race condition

A race condition is when two or more routines have access to the same resource (like variable 'a' in code [`race.go`]) and attempt to read and write to that resource without any regard to the other routines.

This type of code can create random bugs or wrong result which is not expected from the code. Also, we can say that the code outcome depend on interleavings so the output will defer all the time.

### How it can occur?

This occurs due to communication between `goroutines`. When `goroutines` communicate through shared variable (like variable 'a' in code [`race.go`]) the race condition occurs.

In the above code [`race.go`] their are 2 `goroutines`. One of the `goroutines` calls function which will increase the value for variable a by 1 and another function will decrease the value of variable a by 1. When the code is executed it prints different output all the time.

The output shown above differs from each other as the interleavings differs i.e, we can not tell which function will be completed 1st. As their can be many interleavings the output will depend on it and output always differs.
