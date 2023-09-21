package main

import (
	"fmt"
)

func Merchant(perso *Character) {
	menumerchant := 0
	fmt.Println("Bienvenue aventurier")
	fmt.Println("1- Acheter une potion de soin (gratuitement)")
	fmt.Println("2- Acheter une potion de poison (gratuitement)")
	fmt.Println("3- Acheter un livre de sort = boule de feu")
	fmt.Println("4- Revenir dans le menu principal")
	fmt.Scan(&menumerchant)

	switch menumerchant { // Menu of merchant in switch case form, permit to execute functions
	case 1:
		AddInventory(perso, "potion de soin")
		perso.money -= 3
		Menu(perso)
	case 2:
		AddInventory(perso, "potion de poison")
		perso.money -= 7
		Menu(perso)
	case 3:
		AddInventory(perso, "Livre de sort: boule de feu")
		perso.money -= 25
		Menu(perso)
	case 4:
		AddInventory(perso, "fourrure de loup")
		perso.money -= 4
		Menu(perso)
	case 5:
		AddInventory(perso, "peau de Troll")
		perso.money -= 7
		Menu(perso)
	case 6:
		AddInventory(perso, "cuir de sanglier")
		perso.money -= 3
		Menu(perso)
	case 7:
		AddInventory(perso, "plume de Corbeau")
		perso.money -= 1
		Menu(perso)
	case 8:
		Menu(perso)
	}
}