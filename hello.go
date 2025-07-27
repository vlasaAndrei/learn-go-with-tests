package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bounjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
