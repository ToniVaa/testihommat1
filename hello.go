package main

import "fmt"


const spanish = "Spanish"
const french = "French"
const finnish = "Finnish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const finnishHelloPrefix = "Moi, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case finnish:
		prefix = finnishHelloPrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello("Toni", "Finnish"))
}
