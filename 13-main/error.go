package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")

	defer file.Close()

	if err != nil {
		fmt.Println("Error al abrir archivo")
		os.Exit(1)
	} else {
		fmt.Println("Se Abrio el Archivo...")
	}
	examplePanic()
}

func examplePanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Error encontro un Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Found Value 1")
	}
}
