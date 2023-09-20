package main

type skill int

const (
    boule_de_feu skill = iota
    éclaire_de_givre skill = iota
    Projectil_des_arcanes skill = iota
)

type Character struct { // Structure of Character
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	arrows      int
	inventory   []string
	skill       []string
}
func Char() Character { // Function that create character with his specifications
	var perso Character
		perso.name = "Legolas"
		perso.race = "elf"
		perso.class = "mage"
		perso.level = 1
		perso.maxlife = 100
		perso.currentlife = 50
		perso.arrows = 50
		perso.inventory = []string{"potion de soin"}
		perso.skill = []string Character.skill
	return perso
}


func main() { // Implementation of character in "perso" variable and start program with Menu function
	perso := Char()
	Menu(&perso)
}
