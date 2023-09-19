package main

import (
	"fmt"
	"os"
)

func Menu(persovar *Character) {
	menuinput := 0
	fmt.Println("Menu:")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println('\n')
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println('\n')
	fmt.Println("3- Utiliser une potion de soin")
	fmt.Println('\n')
	fmt.Println("4- Quitter")
	fmt.Println('\n')
	fmt.Println("Votre choix ?")
	fmt.Scan(&menuinput)

	switch menuinput {
	case 1:
		DisplayInfo(persovar)
	case 2:
		os.Exit(0)
		///AccessInventory()
	case 3:
		os.Exit(0)
	case 4:
		os.Exit(0)
	}
}
