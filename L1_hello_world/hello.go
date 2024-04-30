package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	korea   = "Korea"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	koreaHelloPrefix   = "Annyeongasheo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case korea:
		prefix = koreaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Ramos", spanish))
	fmt.Println(Hello("Mbappe", french))
	fmt.Println(Hello("Rose", korea))
	fmt.Println(Hello("", ""))
}
