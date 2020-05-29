package main

import "fmt"

var result int

func main() {
	fmt.Println("Start")

	result = operationsMidd(add)(2, 3)
	fmt.Println(result)
	result = operationsMidd(res)(10, 5)
	fmt.Println(result)
	result = operationsMidd(multi)(3, 3)
	fmt.Println(result)
}

func add(a int, b int) int {
	return a + b
}

func multi(a int, b int) int {
	return a * b
}

func res(a int, b int) int {
	return a - b
}

func operationsMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Start Operation")
		return f(a, b)
	}
}
