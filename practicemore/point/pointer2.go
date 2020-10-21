package main

import "fmt"

func main() {
	toChange := "hello"
	var pointer *string = &toChange
	fmt.Println(*pointer, &pointer)
}
