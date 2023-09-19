package main

import (
	"fmt"
	"os"
)

func Menu() {
	menu := 0
	fmt.Println("Menu:")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println('\n')
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println('\n')
	fmt.Println("3- Utiliser une potion de soin")
	fmt.Println('\n')
	fmt.Println("4- Quitter")

	switch menu {
	case 1:
		DisplayInfo(perso1)
	case 2:
		AccessInventory(perso1)
	case 3:
		TakePot()
	case 4:
		os.Exit(0)
	}
}
