package main

import "fmt"

func main() {
	//Array
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 87
	fmt.Println(a, b)

	//Slice
	c := []int{1, 2, 3}
	d := c
	fmt.Println(c, d)
	c[1] = 87
	fmt.Println(c, d)
}
