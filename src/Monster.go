package main

type Monster struct { // La structure du monstre avec ses caractéristiques
	name        string
	maxlife     int
	currentlife int
	attackpoint int
	exppoint	int
}

func (m *Monster) Init(name string, maxlife int, currentlife int, attackpoint int, exppoint int) {
	m.name = name
	m.maxlife = maxlife
	m.currentlife = currentlife
	m.attackpoint = attackpoint
	m.exppoint = exppoint
}