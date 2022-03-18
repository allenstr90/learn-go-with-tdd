package main

import "fmt"

const helloPrefix = "hello "
const holaPrefix = "hola "
const bonjourPrefix = "bonjour "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == "es" {
		return holaPrefix + name + " !!!"
	}

	return greetingPrefix(lang) + name + " !!!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "es":
		prefix = holaPrefix
	case "fr":
		prefix = bonjourPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Print(Hello("world", "en"))
}
