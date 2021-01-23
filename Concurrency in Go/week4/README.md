# Module 3 Activity - SYNCHRONIZED COMMUNICATION

[Assignment](https://www.coursera.org/learn/golang-concurrency/peer/MAemV/module-4-activity)

## Instructions

The goals of this activity are to explore the design of concurrent algorithms and resulting synchronization issues.

Implement the dining philosopher's problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a **host** which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints "starting to eat \<number\>" on a line by itself, where \<number\> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints "finishing eating \<number\>" on a line by itself, where \<number\> is the number of the philosopher.

## My assignment

Source code at `philosopher.go`

## Testing

```text
go run philosopher.go

-----------------------------------------------------------------------------------------------------
MISSION BEGINS
------------------------------------------------------------------------------------------------------
starting to eat  <3>            ( Using ChopStickticks: #2 and #3 )
starting to eat  <1>            ( Using ChopStickticks: #0 and #1 )
finishing eating <3>            ( Eating turns consumed: 1. Units of time required: 815 )
finishing eating <1>            ( Eating turns consumed: 1. Units of time required: 928 )
------------------------------------------------------------------------------------------------------
starting to eat  <4>            ( Using ChopStickticks: #3 and #4 )
starting to eat  <2>            ( Using ChopStickticks: #1 and #2 )
finishing eating <2>            ( Eating turns consumed: 1. Units of time required: 302 )
finishing eating <4>            ( Eating turns consumed: 1. Units of time required: 587 )
------------------------------------------------------------------------------------------------------
starting to eat  <5>            ( Using ChopStickticks: #4 and #0 )
starting to eat  <3>            ( Using ChopStickticks: #2 and #3 )
finishing eating <5>            ( Eating turns consumed: 1. Units of time required: 160 )
finishing eating <3>            ( Eating turns consumed: 2. Units of time required: 815 )
------------------------------------------------------------------------------------------------------
starting to eat  <1>            ( Using ChopStickticks: #0 and #1 )
starting to eat  <4>            ( Using ChopStickticks: #3 and #4 )
finishing eating <4>            ( Eating turns consumed: 2. Units of time required: 587 )
finishing eating <1>            ( Eating turns consumed: 2. Units of time required: 928 )
------------------------------------------------------------------------------------------------------
starting to eat  <2>            ( Using ChopStickticks: #1 and #2 )
starting to eat  <5>            ( Using ChopStickticks: #4 and #0 )
finishing eating <5>            ( Eating turns consumed: 2. Units of time required: 160 )
finishing eating <2>            ( Eating turns consumed: 2. Units of time required: 302 )
------------------------------------------------------------------------------------------------------
starting to eat  <3>            ( Using ChopStickticks: #2 and #3 )
starting to eat  <1>            ( Using ChopStickticks: #0 and #1 )
finishing eating <3>            ( Eating turns consumed: 3. Units of time required: 815 )
finishing eating <1>            ( Eating turns consumed: 3. Units of time required: 928 )
------------------------------------------------------------------------------------------------------
starting to eat  <4>            ( Using ChopStickticks: #3 and #4 )
starting to eat  <2>            ( Using ChopStickticks: #1 and #2 )
finishing eating <2>            ( Eating turns consumed: 3. Units of time required: 302 )
finishing eating <4>            ( Eating turns consumed: 3. Units of time required: 587 )
------------------------------------------------------------------------------------------------------
starting to eat  <5>            ( Using ChopStickticks: #4 and #0 )
finishing eating <5>            ( Eating turns consumed: 3. Units of time required: 160 )
------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------
MISSION ACCOMPLISHED
------------------------------------------------------------------------------------------------------
```
