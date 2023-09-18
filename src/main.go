package main

import (
	"fmt"
)

type Personne struct {
	race string
	class string
	life int
	mana int
	energy int
}

func main() {
	var perso1 Personne

	perso1.race = "elf"
	fmt.Println(perso1.race)

}