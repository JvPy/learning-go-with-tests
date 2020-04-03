package main

import "fmt"

const englishHello = "Hello, "
const spanishHello = "Hola, "
const portugueseHello = "Olá, "

func main() {
	fmt.Println(SayHello("Joao", "Portuguese"))
}

func SayHello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case "Portuguese":
		prefix = portugueseHello
	case "Spanish":
		prefix = spanishHello
	default:
		prefix = englishHello
	}

	return
}
