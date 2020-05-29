package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.User
}

func main() {
	u := new(pepe)
	u.SetupUser(1, "Camilo", time.Now(), true)
	fmt.Println(u.User)
}
