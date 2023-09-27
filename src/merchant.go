package main

import (
	"fmt"
)

func Merchant(perso *Character, monster *Monster) { // Afficher les options de l'interface marchant
	menumerchant := 0
	Clear()
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

	switch menumerchant { // Code pour ajouter l'item dans l'inventaire et retirer l'argent au joueur
	case 1:
		AddInventory(perso, "potion de soin")
		perso.money -= 3
		Merchant(perso, monster)
	case 2:
		AddInventory(perso, "potion de poison")
		perso.money -= 7
		Merchant(perso, monster)
	case 3:
		AddInventory(perso, "Livre de sort: boule de feu")
		perso.money -= 25
		Merchant(perso, monster)
	case 4:
		AddInventory(perso, "fourrure de loup")
		perso.money -= 4
		Merchant(perso, monster)
	case 5:
		AddInventory(perso, "peau de troll")
		perso.money -= 7
		Merchant(perso, monster)
	case 6:
		AddInventory(perso, "cuir de sanglier")
		perso.money -= 3
		Merchant(perso, monster)
	case 7:
		AddInventory(perso, "plume de corbeau")
		perso.money -= 30
		Merchant(perso, monster)
	case 8:
		UpgradeInventorySlot(perso)
		perso.money -= 30
		Merchant(perso, monster)
	case 0:
		Menu(perso, monster)
	}
}