package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go NameSlowy("Camilo Araque")
	var x string
	fmt.Println("Escibe alguna cosa...:")
	fmt.Scanln(&x)
}

func NameSlowy(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%s", letter)
	}
}
