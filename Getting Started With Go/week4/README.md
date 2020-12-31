# Module 4 Activity

[Assignment - makejson](https://www.coursera.org/learn/golang-getting-started/peer/0xyF8/module-4-activity-makejson-go)
[Assignment - read](https://www.coursera.org/learn/golang-getting-started/peer/CHgQd/final-course-activity-read-go)

## Instructions - makejson

The goal of this program is to practice working with RFCs in Go, using JSON as an example.

Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

## My assignment - makejson

Source code at `makejson.go`

## Testing - makejson

```text
go run makejson.go
Enter a name, X to quit: Earl
Enter an address, X to quit: 123 Main Street
{
   "address": "123 Main Street",
   "name": "Earl"
}
```

## Instructions - read

The goal of this assignment is to practice working with the ioutil and os packages in Go.

Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

## My assignment - read

Source code at `read.go`

## Testing - read

```text
go run read.go
Enter the name of data file: data.txt
Edith Whiteaker
John Parks
Harvey Freeman
Jay Wells
David Bowman
Robert Henkle
Janice Buggs
Richard Tolentino
Sarrrrrrrrrrrrrrrrrr Smmmmmmmmmmmmmmmmmmm
Isaac Davis
Johanna Hasler
Jacques Raynor
Buster Kirlin
Madison Jakubowski
Taryn DuBuque
Eryn Schoen
Tørjus Thu
Sarah Steuber
Saaaaaaaaaaaaaaaaaaa Jooooooooooooooooooo```
```

## Input data file

```text
Edith Whiteaker
John Parks
Harvey Freeman
Jay Wells
David Bowman
Robert Henkle
Janice Buggs
Richard Tolentino
Sarrrrrrrrrrrrrrrrrrrrrrrrrah Smmmmmmmmmmmmmmmmmmmmmmith
Isaac Davis
Johanna Hasler
Jacques Raynor
Buster Kirlin
Madison Jakubowski
Taryn DuBuque
Eryn Schoen
Tørjus Thu
Sarah Steuber
Saaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaam Jooooooooooooooooooooooooones
```
