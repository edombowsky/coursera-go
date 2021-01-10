package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a sequence of integers (more than 4), X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			break
		}

		if len(text) != 0 {
			v := strings.Fields(text)
			// if len(v) > 10 {
			// 	v = v[0:10]
			// }

			values := []int{}
			for i := range v {
				number := v[i]
				num, _ := strconv.Atoi(number)
				values = append(values, num)
			}

			numberOfIntegers := len(values)

			if numberOfIntegers < 4 {
				fmt.Println("Enter more than 4 integers, try again")
				continue
			}

			// dividing the slice in 4 equal parts
			numberIntegersEachSlice := numberOfIntegers / 4
			// divLen := len(s) / 4
			reminder := numberOfIntegers % 4
			var slice1, slice2, slice3, slice4 []int
			var lens1, lens2, lens3 int

			// Finding the lengths of each sub slice according to reminder
			switch reminder {
			case 0:
				lens1 = numberIntegersEachSlice
				lens2 = numberIntegersEachSlice * 2
				lens3 = numberIntegersEachSlice * 3
			case 1:
				lens1 = numberIntegersEachSlice + 1
				lens2 = numberIntegersEachSlice*2 + 1
				lens3 = numberIntegersEachSlice*3 + 1
			case 2:
				lens1 = numberIntegersEachSlice + 1
				lens2 = numberIntegersEachSlice*2 + 2
				lens3 = numberIntegersEachSlice*3 + 2
			case 3:
				lens1 = numberIntegersEachSlice + 1
				lens2 = numberIntegersEachSlice*2 + 2
				lens3 = numberIntegersEachSlice*3 + 3
			default:
				fmt.Println("Should never happen, lol. Start over.")
				continue
			}
			// if rem == 0 {
			// 	lens1 = divLen
			// 	lens2 = divLen * 2
			// 	lens3 = divLen * 3
			// } else if rem == 1 {
			// 	lens1 = divLen + 1
			// 	lens2 = divLen*2 + 1
			// 	lens3 = divLen*3 + 1
			// } else if rem == 2 {
			// 	lens1 = divLen + 1
			// 	lens2 = divLen*2 + 2
			// 	lens3 = divLen*3 + 2
			// } else if rem == 3 {
			// 	lens1 = divLen + 1
			// 	lens2 = divLen*2 + 2
			// 	lens3 = divLen*3 + 3
			// }

			//creating 4 sub slice according to the length obtained from above
			slice1 = values[:lens1]
			slice2 = values[lens1:lens2]
			slice3 = values[lens2:lens3]
			slice4 = values[lens3:]

			fmt.Println("\n--- Unsorted --- \n\n", values)

			fmt.Printf("\nTotal number of integers entered [%d]\n", numberOfIntegers)
			fmt.Printf("Number of integers in each slice [%d]\n", numberIntegersEachSlice)

			fmt.Printf("\nInitial slice of integers to be sorted in Thread #%d: %v\n", 1, slice1)
			fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 2, slice2)
			fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 3, slice3)
			fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 4, slice4)

			var wg sync.WaitGroup
			wg.Add(4)
			go BubbleSort(&wg, 1, slice1)
			go BubbleSort(&wg, 2, slice2)
			go BubbleSort(&wg, 3, slice3)
			go BubbleSort(&wg, 4, slice4)
			wg.Wait()

			fmt.Printf("\nResulting slice of integers sorted in Thread #%d: %v\n", 1, slice1)
			fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 2, slice2)
			fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 3, slice3)
			fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 4, slice4)

			// Merging the slice sorted sub slice into one slice
			mergedSlice := make([]int, numberOfIntegers)
			for i, v := range slice1 {
				mergedSlice[i] = int(v)
			}
			for i, v := range slice2 {
				mergedSlice[i+len(slice1)] = int(v)
			}
			for i, v := range slice3 {
				mergedSlice[i+len(slice1)+len(slice2)] = int(v)
			}
			for i, v := range slice4 {
				mergedSlice[i+len(slice1)+len(slice2)+len(slice3)] = int(v)
			}
			fmt.Printf("\nAfter merging all that slices, the initial slice of integers now contains: %v\n", mergedSlice)

			sort.Ints(mergedSlice)

			fmt.Println("\n--- Sorted ---\n\n", mergedSlice)
		}
	}
}

// BubbleSort will sort an integer slice
func BubbleSort(wg *sync.WaitGroup, threadID int, items []int) {

	fmt.Printf("On thread #[%d] - Slice received %v\n", threadID, items)
	fmt.Printf("On thread #[%d] - Sorting ...\n", threadID)

	sort.Ints(items)

	fmt.Printf("On thread #[%d] - Slice sorted %v. Exiting thread...\n", threadID, items)
	wg.Done()
}
