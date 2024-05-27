package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	greeting := "Hello, OTUS!"
	reversedGreeting := reverse.String(greeting)
	fmt.Println(reversedGreeting)
}
