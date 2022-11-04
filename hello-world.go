package main

import (
	"fmt"
)

func printer(data string) {
	fmt.Println(data)
}

func loopy() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

func loopy2(max int) {
	i := 1
	for i < max {
		fmt.Println(i)
		i = i + 1
	}
}

func loopy3() {
	for i := 3; i < 20; i++ {
		fmt.Print("shit ")
	}
}

func fizzbuzz(max int) {
	for i := 1; i < max; i++ {
		output := ""
		if i%3 == 0 && i%5 == 0 {
			output = output + "FizzBuzz"
		} else if i%3 == 0 {
			output = output + "Fizz"
		} else if i%5 == 0 {
			output = output + "Buzz"
		} else {
			fmt.Print(i)
		}
		fmt.Println(output)
	}
}

func main() {
	const constant_value string = "This is a constant string"

	// message1 := fmt.Sprint("Hello")
	// message0 := fmt.Sprint(123)

	// fmt.Println(message1, message0)

	// printer(constant_value)

	// loopy()
	// loopy2(12)
	// loopy3()

	fizzbuzz(21)

}
