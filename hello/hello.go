package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const lithuanian = "Lithuanian"
const greetingEnglish = "Hello, "
const greetingSpanish = "Hola, "
const greetingFrench = "Bonjour, "
const greetingLithuanian = "Sveiki, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingEnglish

	switch language {
	case french:
		prefix = greetingFrench
	case spanish:
		prefix = greetingSpanish
	case lithuanian:
		prefix = greetingLithuanian
	}

	return prefix + name
	}

func main() {
	fmt.Println(Hello("world", ""))
}
