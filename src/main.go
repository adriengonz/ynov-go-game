package main

import (
)

type Personne struct {
	name string
	race string
	class  string
	level int
	maxlife   int
	currentlife  int
	inventory []string
}

func GetInfoChar() {
	var perso1 Personne
	perso1.name = "Legolas"
	perso1.race = "elf"
	perso1.class = "archer"
	perso1.level = 1
	perso1.maxlife = 100
	perso1.currentlife = 50
	perso1.inventory = [] string{"3 potion de soin", "50 arrows"}
}

func main() {
	Menu()
}