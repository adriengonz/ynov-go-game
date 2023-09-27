package main

type Monster struct { // La structure du monstre avec ses caractéristiques
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