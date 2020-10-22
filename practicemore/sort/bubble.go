package main

import "fmt"

var toBeSorted [10]int = [10]int{1, 5, 6, 3, 2, 7, 9, 6, 4, 1}

func bubblesort(input [10]int) {
	n := 10
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				//fmt.Println("Swapping..!")
				input[i-1], input[i] = input[i], input[i-1]
				//fmt.Println(input)
				swapped = true
			}
		}
	}
	fmt.Println(input)
}
func main() {
	fmt.Println(toBeSorted)
	bubblesort(toBeSorted)
}
