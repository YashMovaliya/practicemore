package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}
