package main

import (
	"fmt"
)

func Merchant(perso *Character) {
	menumerchant := 0
	fmt.Println("Bienvenue aventurier")
	fmt.Println("1- Acheter une potion de soin (Coût : 3 pièces d'or)")
	fmt.Println("2- Acheter une potion de poison (Coût : 7 pièces d'or)")
	fmt.Println("3- Acheter un livre de sort : boule de feu (Coût : 25 pièces d'or)")
	fmt.Println("4- Acheter une fourrure de loup (Coût : 4 pièces d'or)")
	fmt.Println("5- Acheter une peau de troll (Coût : 7 pièces d'or)")
	fmt.Println("6- Acheter du cuir de sanglier (Coût : 3 pièces d'or)")
	fmt.Println("7- Acheter une plume de corbeau (Coût : 1 pièce d'or)")
	fmt.Println("8- Acheter 10 emplacements de sac (Coût : 30 pièce d'or)")
	fmt.Println("0- Revenir dans le menu principal")
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
		AddInventory(perso, "peau de troll")
		perso.money -= 7
		Menu(perso)
	case 6:
		AddInventory(perso, "cuir de sanglier")
		perso.money -= 3
		Menu(perso)
	case 7:
		AddInventory(perso, "plume de corbeau")
		perso.money -= 30
		Menu(perso)
	case 8:
		UpgradeInventorySlot(perso)
		perso.money -= 30
		Menu(perso)
	case 9:
		Menu(perso)
	}
}
