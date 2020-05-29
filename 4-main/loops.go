package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n Valor de i es: %d", i)
		if i == 5 {
			fmt.Printf("\t Number miltiplicated for 2")
			i = i * 2
			continue
		}
		fmt.Printf("\t Pase Por aquí")
	}
	rutine()
}

func rutine() {
	var num int
RUTINE:
	for num <= 10 {
		if num == 4 {
			fmt.Printf("\n Sume la Iteración en 3 a num")
			num = num + 3
			goto RUTINE
		}
		fmt.Printf("\n Valor de i: %d", num)
		num++
	}
}
