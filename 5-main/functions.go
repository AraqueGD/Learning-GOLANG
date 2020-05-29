package main

import "fmt"

func main() {
	fmt.Println(one(5))
	rta, status := two(1)
	fmt.Println(rta, status)
	fmt.Println("Sum Total: ", three(5, 46))
	fmt.Println("Sum Total: ", three(5, 46, 56))
	fmt.Println("Sum Total: ", three(1, 67, 7, 78))
}

func one(number int) int {
	return number * 2
}

func two(number int) (int, bool) {
	if number == 1 {
		return number * 2, true
	}
	return number * 3, false
}

func three(number ...int) int {
	total := 0
	for item, num := range number {
		total = total + num
		fmt.Printf("You Index is %d: %d\n", item, num)
	}
	return total
}
