package main

type Character struct { // Structure of Character
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	arrows      int
	inventory   []string
}
func Char() Character { // Function that create character with his specifications
	var perso Character
		perso.name = "Legolas"
		perso.race = "elf"
		perso.class = "archer"
		perso.level = 1
		perso.maxlife = 100
		perso.currentlife = 50
		perso.arrows = 50
		perso.inventory = []string{"potion de soin"}
	return perso
}


func main() { // Implementation of character in "perso" variable and start program with Menu function
	perso := Char()
	Menu(&perso)
}
