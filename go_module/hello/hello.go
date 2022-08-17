package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Zsena")
	fmt.Println(message)
}