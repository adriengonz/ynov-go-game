package main

import (
	"fmt"
	"os"
)

func main() { // Fonction pour afficher l'interface du choix de personnage
	var m1 Monster
	m1.Init("Gobelin d’entrainement", 40, 20, 5, 50)
	characterinput := 0
	fmt.Println("Faites votre choix de personnage:")
	fmt.Println("-------------------")
	fmt.Println("1 - Merlius\n","   Human\n","   mage\n","   level 1\n","   pv max = 100\n","   pv actuel = 50")
	fmt.Println("-------------------")
	fmt.Println("2 - Legolas\n","   Elf\n","   archer\n","   level 1\n","   pv max = 80\n","   pv actuel = 40")
	fmt.Println("-------------------")
	fmt.Println("3 - Baldur\n","   Dwarf\n","   war\n","   level 1\n","   pv max = 120\n","   pv actuel = 60")
	var p1 Character
	fmt.Scan(&characterinput)

	switch characterinput { // Menu pour sélectionner le perso voulu avec les stats et les caractéristiques
	case 1:
		p1.Init("Merlius", "Human", "mage", 1, 100, 50, []string {"potion de soin"}, []string {"coup de poing"}, 100, p1.equipment, 10, 30, 0, 1)
	case 2:
		p1.Init("Legolas", "Elf", "archer", 1, 80, 40, []string {"potion de soin"}, []string {"coup de poing"}, 100, p1.equipment, 10, 30, 0, 1)
	case 3:
		p1.Init("Baldur", "Dwarf", "war", 1, 120, 60, []string {"potion de soin"}, []string {"coup de poing"}, 100, p1.equipment, 10, 30, 0, 1)
	case 0:
		os.Exit(0)
	}
	Menu(&p1, &m1)
}