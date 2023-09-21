package main

import "fmt"

type skill int

const (
    boule_de_feu skill = iota
    éclaire_de_givre
    Projectil_des_arcanes
)

type name int

const (
    Humain skill = iota
    Elf
    Nain
)

type Character struct { // Structure of Character
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	arrow		int
	inventory   []string
	skill       []string
}
func Char() Character { // Function that create character with his specifications
	var perso Character
	fmt.Println("Quel est votre nom ?")
	fmt.Scan(perso.name)
	fmt.Println("Quel est votre race ?")
	fmt.Scan(perso.race)
	fmt.Println("Quel est votre classe ?")
	fmt.Scan(perso.class)
		perso.level = 1
		perso.maxlife = 100
		perso.currentlife = 50
		perso.arrow = 50
		perso.inventory = []string{"potion de soin"}
		perso.skill = []string {"boule de feu", "/" ,"éclaire de givre", "/" , "Projectil des arcanes"}
	return perso
}


func main() { // Implementation of character in "perso" variable and start program with Menu function
	perso := Char()
	Menu(&perso)
}
