package main

func main() {
	var p1 Character
	p1.Init("Legolas", "elf", "mage", 1, 100, 50, []string {"potion de soin"}, []string {"éclaire de givre"}, 100, p1.equipment, 10)
	Menu(&p1)
}

func InitGobelin() {
	var m1 Monster
	m1.Init("Gobelin d’entrainement", 40, 20, 5)
}