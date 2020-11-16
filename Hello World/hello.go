package main

import "fmt"

//package GO_Tutorial_Learn_GO_with_Test
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Wolle"))
}
