package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
	readFile2()
	writeFile()
	writeFile2()
}

func readFile() {
	file, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./text.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			register := scanner.Text()
			fmt.Println(register)
		}
	}
	file.Close()
}

func writeFile() {
	file, err := os.Create("./NewFile.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
		return
	}
	fmt.Fprintln(file, "This is a new text line")
	file.Close()
}
func writeFile2() {
	file := "./NewFile.txt"
	if Append(file, "\n New LINE YESSS!!!!") == false {
		fmt.Println("Hubo un Error")
		return
	}
}

func Append(file string, text string) bool {
	fileName, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un Error")
		return false
	}
	_, err = fileName.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un Error")
		return false
	}
	return true
}
