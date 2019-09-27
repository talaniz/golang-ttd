package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

// Hello returns a greeting
func Hello(name string) string {
	return englishHelloPrefix + name
}
