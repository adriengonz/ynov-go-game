package main

func main() {
	var p1 Character
	p1.Init("Legolas", "elf", "mage", 1, 100, 50, []string {"potion de soin"}, []string {"éclaire de givre"}, 100, p1.equipment, 10)
	var m1 Monster
	m1.Init("Gobelin d’entrainement", 40, 20, 5)
	Menu(&p1, &m1)
}