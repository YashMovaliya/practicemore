package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Hello World")
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		fmt.Println("GoodBye World")
	})
	http.ListenAndServe(":8080", nil)
}
