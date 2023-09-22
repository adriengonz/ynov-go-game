package main

type skill int

const (
    boule_de_feu skill = iota
    éclaire_de_givre
    Projectil_des_arcanes
)

type name int

const (
    Humain name = iota
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
	money		int
}

type Equipment struct { // Structure of Equipment
	head 		[]string
	torso       []string
	foot		[]string
}

func Char() Character { // Function that create character with his specifications
	var perso Character
		perso.name = "Legolas"
		perso.race = "elf"
		perso.class = "mage"
		perso.level = 1
		perso.maxlife = 100
		perso.currentlife = 50
		perso.arrow = 50
		perso.inventory = []string{"potion de soin"}
		perso.skill = []string {"boule de feu", "/" ,"éclaire de givre", "/" , "Projectil des arcanes"}
		perso.money = 100
	return perso
}

func stuf() Equipment {
    var stuff Equipment
    stuff.head = []string{"Emplacement du casque"}
    stuff.torso = []string{"Emplacement du torse"}
    stuff.foot = []string{"Emplacement des bottes"}
    return stuff
}

func main() { // Implementation of character in "perso" variable and start program with Menu function
    perso := Char()
    stuff := stuf()
    Menu(&perso)
}