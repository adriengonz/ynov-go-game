package main

import (
	"fmt"
	"main/menu/infocharacter"
)

type Personne struct {
	race   string
	class  string
	life   int
	mana   int
	energy int
}

func main() {

	var perso1 Personne
	fmt.Println(Tkt)
	perso1.race = "elf"
	fmt.Println(perso1.race)

}
