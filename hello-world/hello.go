package main

import "fmt"

const helloPrefix = "hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name + " !!!"
}

func main() {
	fmt.Print(Hello("world"))
}
