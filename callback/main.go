package main

import "fmt"

// process input then output the processed result
type calc func(message string) int

// produce hash from source
func produce(msg int) string {
	if msg < 0 {
		return "negative"
	} else if msg == 0 {
		return "zero"
	}
	return "positive"
}
func consume(msg string, calcter calc) int {
	return calcter(msg)
}
func main() {
	msg := produce(1000)
	fmt.Println("result: ", consume(msg, func(message string) int {
		fmt.Println("process:", msg)
		return 0
	}))
}
