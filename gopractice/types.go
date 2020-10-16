package main

import "fmt"

func main() {
var maxFloat float32
maxFloat = 16777216
fmt.Println(maxFloat == maxFloat+10)

fmt.Println(maxFloat+10)
fmt.Println(maxFloat+200000)
}