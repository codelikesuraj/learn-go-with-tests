package main

import "fmt"

func main() {
	fmt.Println(Hello("world", "lks"))
}

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix string = "Hello, "
const frenchHelloPrefix string = "Bonjour, "
const spanishHelloPrefix string = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixGreeting(language) + name + "!"
}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}