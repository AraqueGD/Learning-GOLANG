package main

import "fmt"

func main() {
	set := make(map[string]string)
	fmt.Println(set)

	set["Colombia"] = "Bogotá D.C"
	set["España"] = "Madrid"

	fmt.Println(set)

	champion := map[string]int{
		"Barcelona":   8,
		"Real Madrid": 3,
		"Chivas":      2,
		"Boca":        6}

	champion["Nacional"] = 10

	delete(champion, "Chivas")

	fmt.Println(champion)

	for team, score := range champion {
		fmt.Printf("The Team %s have score: %d\n", team, score)
	}
	score, exists := champion["Junior"]
	fmt.Printf("The Socore is %d, team is exists: %t\n", score, exists)
}
