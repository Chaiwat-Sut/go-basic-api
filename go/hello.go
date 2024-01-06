package main

import (
	"fmt"
	"strconv"
)

type change struct {
	Banknote int
	Amount   int
}

var banknote = []int{1000, 500, 100, 50, 20, 10, 5, 2, 1}

func sell(price, pay int) []change {
	result := []change{}
	c := pay - price
	for _, note := range banknote {
		if c >= note {
			noteAmount := c / note
			c = c - (noteAmount * note)
			result = append(result, change{Banknote: note, Amount: noteAmount})
		}
	}
	return result
}

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Printf("%v\n", fizzBuzz(1))
	// fmt.Printf("%v\n", fizzBuzz(3))
	// fmt.Printf("%v\n", fizzBuzz(5))
	// fmt.Printf("%v\n", fizzBuzz(7))
	// fmt.Printf("%v\n", fizzBuzz(15))

	fmt.Printf("%v\n", sell(50, 100))
}

func fizzBuzzLoop(numSlice []int) {
	for _, value := range numSlice {
		fmt.Printf("%v\n", fizzBuzz(value))
	}
}

func fizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
