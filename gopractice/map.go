package main

import "fmt"

func main() {
	x := map[string]int{"Kate": 28, "John": 37, "Raj": 20}
	fmt.Print(x)
	fmt.Println("\n", x["Raj"])
}
