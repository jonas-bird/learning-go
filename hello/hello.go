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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = greetingFrench
	case spanish:
		prefix = greetingSpanish
	case lithuanian:
		prefix = greetingLithuanian
	default:
		prefix = greetingEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
