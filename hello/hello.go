package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const greetingEnglish = "Hello, "
const greetingSpanish = "Hola, "
const greetingFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return greetingSpanish + name
	}
	if language == french {
		return greetingFrench + name
	}
	return greetingEnglish + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
