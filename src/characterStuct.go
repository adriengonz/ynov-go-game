package main

type name int

const (
    Humain name = iota
    Elf
    Nain
)

type skill int

const (
    éclaire_de_givre skill = iota
    Projectil_des_arcanes
)

type Equipment struct { // Structure of Equipment
	head  []string
	torso []string
	foot  []string
}

type Character struct { // Structure of Character
	name        string
	race        string
	class       string
	level       int
	maxlife     int
	currentlife int
	inventory   []string
	skill       []string
	money       int
	equipment   Equipment
}

func (p *Character) Init(name string, race string, class string, level int, maxlife int, currentlife int, inventory []string, skill []string, money int, equipment Equipment) {
	p.name = name
	p.race = race
	p.class = class
	p.level = level
	p.maxlife = maxlife
	p.currentlife = currentlife
	p.inventory = inventory
	p.skill = skill
	p.money = money
	p.equipment = equipment
}
