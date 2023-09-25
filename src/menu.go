package main

import (
	"fmt"
	"os"
)

func Menu(persovar *Character, monsterinstance *Monster) { // Function that prints the menu of game, who takes in argument "persovar" in pointer, who appoint the character structure
	menuinput := 0
	Clear()
	fmt.Println("Menu:")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println("3- Marchand")
	fmt.Println("4- Forgeron")
	fmt.Println("5- S'entrainer au combat")
	fmt.Println("0- Quitter")
	fmt.Println("Votre choix ?")
	fmt.Scan(&menuinput)

	switch menuinput { // Menu in switch case form, permit to execute functions
	case 1:
		DisplayInfo(persovar, monsterinstance)
	case 2:
		AccessInventory(persovar, monsterinstance)
	case 3:
		Merchant(persovar, monsterinstance)
	case 4:
		Blacksmith(persovar, monsterinstance)
	case 5:
		trainingFight(persovar, monsterinstance)
	case 0:
		os.Exit(0)
	}
}

func Clear(){
	fmt.Print("\033[H\033[2J")
}