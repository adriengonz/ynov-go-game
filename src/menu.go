package main

import (
	"fmt"
	"os"
)

func Menu(persovar *Character) { // Function that prints the menu of game, who takes in argument "persovar" in pointer, who appoint the character structure
	menuinput := 0
	fmt.Println("Menu:")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println("3- Utiliser une potion de soin")
	fmt.Println("4- Marchand")
	fmt.Println("5- Quitter")
	fmt.Println("Votre choix ?")
	fmt.Scan(&menuinput)

	switch menuinput { // Menu in switch case form, permit to execute functions
	case 1:
		DisplayInfo(persovar)
	case 2:
		AccessInventory(persovar)
	case 3:
		TakePot(persovar)
	case 4:
		Merchant(persovar)
	case 5:
		os.Exit(0)
	}
}
