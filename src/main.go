package main

func main() {
	var p1 Character
	p1.Init("Legolas", "elf", "mage", 1, 100, 50, []string {"healing potion"}, []string {"Frost Bolt"}, 100)
	Menu(&p1)
}