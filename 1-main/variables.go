package main

import (
	"fmt"
)

var num int
var text string
var status bool

func main() {

	numero, numero2, numero3, text, status := 4, 5, 8, "Go is perfect!", true

	var moneda int = 3

	text = fmt.Sprintf("%d", moneda)

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(text)
	fmt.Println(status)
	mostrarStatus()
}
func mostrarStatus() {
	fmt.Println(status)
}
