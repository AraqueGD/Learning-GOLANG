package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number int
	var number2 int
	var result int
	var leyend string
	fmt.Println("Enter Number: ")
	fmt.Scanln(&number)

	fmt.Println("Enter Number 2: ")
	fmt.Scanln(&number2)

	fmt.Println(number + number2)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyend = scanner.Text()
	}
	result = number + number2
	fmt.Println(leyend, result)
}
