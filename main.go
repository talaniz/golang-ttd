package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHellloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", "English"))
}

// Hello returns a greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishHellloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}
