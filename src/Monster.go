package main

type Monster struct { // Structure of Character
	name        string
	maxlife     int
	currentlife int
	attackpoint int
}

func (m *Monster) Init(name string, maxlife int, currentlife int, attackpoint int) {
	m.name = name
	m.maxlife = maxlife
	m.currentlife = currentlife
	m.attackpoint = attackpoint
}