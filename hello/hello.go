package main

import "fmt"

const spanish = "Spanish"
const greetingEnglish = "Hello, "
const greetingSpanish = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return greetingSpanish + name
	}
	return greetingEnglish + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
