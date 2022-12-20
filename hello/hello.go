package main

import "fmt"

const greetingEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingEnglish + name
}

func main() {
	fmt.Println(Hello("world"))
}
