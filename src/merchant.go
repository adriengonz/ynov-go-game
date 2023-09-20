package main

import (
	"fmt"
	"os"
)

func Merchant(perso *Character) {
	menuinput := 0
	fmt.Println("Bienvenue aventurier")
	fmt.Println("1- Acheter une potion de soin (gratuitement)")
	fmt.Println("2- Acheter une potion de poison (gratuitement)")
	fmt.Println("3- Acheter un livre de sort = boule de feu")
	fmt.Println("4- Revenire dans le menu principale")
	fmt.Scan(&menuinput)

	switch menuinput { // Menu in switch case form, permit to execute functions
	case 1:
		AddInventory(perso, "potion de soin")
	case 2:
		//(perso, "potion de poison")
	case 3:
		//(perso, "Livre de sort: boule de feu")
	case 4:
		os.Exit(0)
	}
	fmt.Println("Potion de soin à bien était ajoutée à votre inventaire")
}