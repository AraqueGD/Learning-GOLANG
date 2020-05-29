package main

import "fmt"

var calc func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sum 5 + 5 = %d", calc(5, 5))

	calc = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("\nRes 5 - 5 = %d", calc(5, 5))

	calc = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("\nRes 5 / 5 = %d", calc(10, 2))

	tabTable := 2
	myTable := table(tabTable)
	for i := 1; i < 11; i++ {
		fmt.Println(myTable())
	}
}

func table(val int) func() int {
	num := val
	secuence := 0
	return func() int {
		secuence++
		return num * secuence
	}
}
