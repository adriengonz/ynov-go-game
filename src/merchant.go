package main

import (
	"fmt"
)

func Merchant(perso *Character) {
	menumerchant := 0
	fmt.Println("Welcome adventurer")
	fmt.Println("1- Buy a healing potion (Cost: 3 gold)")
	fmt.Println("2- Buy a potion of poison (Cost: 7 gold)")
	fmt.Println("3- Buy a spell book: fireball (Cost: 25 gold)")
	fmt.Println("4- Buy a wolf fur (Cost: 4 gold)")
	fmt.Println("5- Buy a troll skin (Cost: 7 gold)")
	fmt.Println("6- Buy Boar Leather (Cost: 3 gold)")
	fmt.Println("7- Buy a Raven Feather (Cost: 1 gold)")
	fmt.Println("8- Return to the main menu")
	fmt.Scan(&menumerchant)

	switch menumerchant { // Menu of merchant in switch case form, permit to execute functions
	case 1:
		AddInventory(perso, "healing potion")
		perso.money -= 3
		Menu(perso)
	case 2:
		AddInventory(perso, "potion of poison")
		perso.money -= 7
		Menu(perso)
	case 3:
		AddInventory(perso, "Spell Book: Fireball")
		perso.money -= 25
		Menu(perso)
	case 4:
		AddInventory(perso, "wolf fur")
		perso.money -= 4
		Menu(perso)
	case 5:
		AddInventory(perso, "troll skin")
		perso.money -= 7
		Menu(perso)
	case 6:
		AddInventory(perso, "wild boar leather")
		perso.money -= 3
		Menu(perso)
	case 7:
		AddInventory(perso, "raven feather")
		perso.money -= 1
		Menu(perso)
	case 8:
		Menu(perso)
	}
}
