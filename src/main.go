package main

import (
	"os"
	"fmt"
)

func main() {
	var m1 Monster
	m1.Init("Gobelin d’entrainement", 40, 20, 5)
	characterinput := 0
	fmt.Println("Faite votre choix de personnage:")
	fmt.Println("1 - Merlius","Human","mage",1,100,50)
	fmt.Println("2 - Legolas","Elf","archer",1,80,40)
	fmt.Println("3 - Baldur","Dwarf","war",1,100,50)
	var p1 Character
	fmt.Scan(&characterinput)

	switch characterinput {
	case 1:
		p1.Init("Merlius", "Human", "mage", 1, 100, 50, []string {"potion de soin"}, []string {"éclaire de givre"}, 100, p1.equipment, 10)
	case 2:
		p1.Init("Legolas", "Elf", "archer", 1, 80, 40, []string {"potion de soin"}, []string {"éclaire de givre"}, 100, p1.equipment, 10)
	case 3:
		p1.Init("Baldur", "Dwarf", "war", 1, 100, 50, []string {"potion de soin"}, []string {"éclaire de givre"}, 100, p1.equipment, 10)
	case 0:
		os.Exit(0)
	}
	Menu(&p1, &m1)
}