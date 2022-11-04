package main

import (
	"fmt"
)

func printer(data string) {
	fmt.Println(data)
}

func main() {
	const constant_value string = "This is a constant string"

	message1 := fmt.Sprint("Hello")
	message0 := fmt.Sprint(123)

	fmt.Println(message1, message0)

	printer(constant_value)

}
