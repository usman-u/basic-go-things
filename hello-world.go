package main

import (
	"fmt"
)

func main() {
	const constant_value string = "This is a constant string"

	message1 := fmt.Sprint("Hello")
	message0 := fmt.Sprint(123)

	fmt.Print(message1+"da"+" "+message0, " ", constant_value)
}
