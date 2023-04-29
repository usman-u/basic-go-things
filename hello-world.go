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

func fizzbuzz(high int) {
	for i := 1; i < high; i++ {
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

func loop4(newi string) {
	// for i := 0; i < max; i++ {
	// 	fmt.Println(i)
	// }
	slice := make([]string, 0)
	// slice[2] = "second"

	slice = append(slice, newi)

	// fmt.Println("2nd Index: ", slice[2])

	// fmt.Println(slice[2])

	for i := 0; i < len(slice); i++ {
		fmt.Println("Index", i, ":", slice[i])
	}
}

func bubble_sort(data []int) {

	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i])
	// }

	swapped := true
	n := len(data)

	for swapped {
		swapped = false
		n = n - 1

		for i := 0; i < n; i++ {
			if data[i] > data[i+1] {
				tmp := data[i]
				data[i] = data[i+1]
				data[i+1] = tmp
				swapped = true
			}
		}

	}
	fmt.Println(data)
}

func insertion_sort(data []string) []string {
	for i := 0; i < len(data); i++ {
		current := data[i]
		index2 := i

		for index2 > 0 && data[index2-1] > current {
			data[index2] = data[index2-1]
			index2 = index2 - 1
		}
		data[index2] = current
	}
	return data
}

func main() {
	// const constant_value string = "This is a constant string"

	// message1 := fmt.Sprint("Hello")
	// message0 := fmt.Sprint(123)

	// fmt.Println(message1, message0)

	// printer(constant_value)

	// loopy()
	// loopy2(12)

	// fizzbuzz(21)

	// test := []int{20, 5, 12, 4, 2, 55, -1}
	// bubble_sort(test)
	// api()

	// test := []int{10, 20, 30, 40, 50, 60, 70}
	data := []string{"Florida", "Georgia", "Delaware", "Alabama"}
	// result := binary_search(test, 30)
	fmt.Println(data)
	result := insertion_sort(data)
	fmt.Println(result)
}
