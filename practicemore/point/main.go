package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) string {
	str = "changed!"
	return str
}
func main() {
	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	//changeValue2(toChange)
	fmt.Println(toChange)
}
