package main

import "fmt"

func main() {
	exponencial(2)
}

func exponencial(num int) {
	if num > 100000000 {
		return
	}
	fmt.Println(num)
	exponencial(num * 2)
}
