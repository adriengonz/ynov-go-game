package main

import "fmt"

type Personne struct {
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	arrows      int
	inventory   []string
}

func GetInfoChar() {
	var perso1 Personne
	perso1.name = "Legolas"
	perso1.race = "elf"
	perso1.class = "archer"
	perso1.level = 1
	perso1.maxlife = 100
	perso1.currentlife = 50
	perso1.arrows = 50
	perso1.inventory = []string{"potion de soin"}
	fmt.Println(perso1.name)
	fmt.Println(perso1.race)
	fmt.Println(perso1.class)
	fmt.Println(perso1.level)
	fmt.Println(perso1.maxlife)
	fmt.Println(perso1.currentlife)
	fmt.Println(perso1.inventory)
}

func main() {
	GetInfoChar()
}
