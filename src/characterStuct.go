package main

type Equipment struct { // La structure de l'équipement
	head  []string
	torso []string
	foot  []string
}

type Character struct { // La structure du personnage
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
	limitInventory int
	mana 		int
	currentExperience int
	turn		int
}

func (p *Character) Init(name string, race string, class string, level int, maxlife int, currentlife int, inventory []string, skill []string, money int, equipment Equipment, limitInventory int, mana int, currentExperience int, turn int) {
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
	p.limitInventory = limitInventory
	p.mana = mana
	p.currentExperience = currentExperience
	p.turn = turn
}