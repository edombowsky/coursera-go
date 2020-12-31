# Module 2 Activity

## trunc.go

[Assignment](https://www.coursera.org/learn/golang-getting-started/peer/01wGe/module-2-activity-trunc-go/submit)

## Instructions - trunc

The goal of this assignment is to practice working with floating point numbers and truncation code in Go.

Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

## My assignment - trunc

Source code at `hello.go`

## Testing - trunc

```text
go run trunc.go
Enter a floating point number, X to quit: 1.2222222
1.222222 truncated = 1
Enter a floating point number, X to quit: 11111111.0989789782893298173891
11111111.098979 truncated = 11111111
Enter a floating point number, X to quit: -103.3344
-103.334400 truncated = -103
Enter a floating point number, X to quit: 
```

## findian.go

[Assignment](https://www.coursera.org/learn/golang-getting-started/peer/f1Cly/module-2-activity-findian-go)

## Instructions - findian

The goal of this assignment is to practice working with strings in Go.

Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print "Found!" if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print "Not Found!" otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

**Examples:** The program should print "Found!" for the following example entered strings, "ian", "Ian", "iuiygaygn", "I d skd a efju N". The program should print "Not Found!" for the following strings, "ihhhhhn", "ina", "xian".

## My assignment - findian

Source code at `findian.go`

## Testing - findian

```text
go run findian.go
Enter a string, X to quit: ian
Found
Enter a string, X to quit: Ian
Found
Enter a string, X to quit: iuiygaygn
Found
Enter a string, X to quit: I d skd a efju N 
Found
Enter a string, X to quit: ihhhhhn
Not Found
Enter a string, X to quit: ina
Not Found
Enter a string, X to quit: xian
Not Found
Enter a string, X to quit: x
```
