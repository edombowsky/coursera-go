# Module 2 Activity

[Assignment](https://www.coursera.org/learn/golang-functions-methods/peer/qKrnv/module-2-activity)

## Instructions

The goal of this assignment is to create a routine that solves a linear kinematics problem.

Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s(t) = ½ (at²) + v0t + s0

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let's say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

## My assignment

Source code at `kinematics.go`

## Testing

```text
go run kinematics.go

Enter acceleration, X to quit: 10
Enter initial velocity, X to quit: 2
Enter initial displacement, X to quit: 1
Enter time, X to quit: 3
The resulting displacement value s(t) is [52.000000]
Enter acceleration, X to quit: 10
Enter initial velocity, X to quit: 2
Enter initial displacement, X to quit: 1
Enter time, X to quit: 5
The resulting displacement value s(t) is [136.000000]
Enter acceleration, X to quit: x
2020/12/26 19:01:08 Scan for acceleration failed, due to strconv.ParseFloat: parsing "": invalid syntax
```
