package main

import "fmt"

const englishHello = "Hello, "

func main() {
	fmt.Println(SayHello("Joao"))
}

func SayHello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHello + name
}
