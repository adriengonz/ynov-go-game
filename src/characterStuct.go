package main

type name int

const (
    Humain name = iota
    Elf
    Nain
)

type skill int

const (
    boule_de_feu skill = iota
    éclaire_de_givre
    Projectil_des_arcanes
)

type Character struct { // Structure of Character
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	arrow       int
	inventory   []string
	skill       []string
	money       int
}

func (p *Character) Init(name string, race string, class string, maxlife int, currentlife int) {
	p.name = name
	p.race = race
	p.class = class
	p.maxlife = maxlife
	p.currentlife = currentlife
}
