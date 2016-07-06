package main

import "fmt"

const (
	FATAL = 50 - 10 * iota
	ERROR
	WARN
	INFO
	DEBUG
)

func main() {
	fmt.Println(FATAL)
	fmt.Println(ERROR)
	fmt.Println(WARN)
	fmt.Println(INFO)
	fmt.Println(DEBUG)
}
