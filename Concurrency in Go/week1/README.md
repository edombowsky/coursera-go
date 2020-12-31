# Module 1 Activity - CONCURRENCY BASICS

[Assignment](https://www.coursera.org/learn/golang-concurrency/peer/RAJ0V/module-2-activity)

## Instructions

The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.

Write two `goroutines` which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

## My assignment

Source code at `animals.go`

## Testing

```text
go run animals.go
> newanimal daisy cow
Created it!
> newanimal big bird
Created it!
> newanimal simon snake
Created it!
> query daisy eat
grass
> query daisy move
walk
> query daisy speak
moo
> query big eat
worms
> query big move
fly
> query big speak
peep
> query simon eat
mice
> query simon move
slither
> query simon speak
hsss
> newanimal felix cat
type of animal [cat] is not in the list of accepted types. Please try again
> query felix eat
there is no such animal [felix] stored. Please try again
> delete big bird
Unknown command 'delete', please try again.
> command name
Bad request, three strings are requested
> x
```
