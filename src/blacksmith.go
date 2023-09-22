package main

import (
	"fmt"
	"time"
)

func Blacksmith(perso *Character) {
	menublacksmith := 0
	fmt.Println("Bienvenue chez le forgeron")
	fmt.Println("1 - Craft an Adventurer's Hat (5 gold, 1 raven feather, and 1 boar hide)")
	fmt.Println("2 - Craft an Adventurer's Tunic (5 gold, 2 wolf fur, and 1 troll skin)")
	fmt.Println("3 - Craft Adventurer's Boots (5 gold, 1 wolf fur, and 1 boar leather)")
	fmt.Println("4 - Return to main menu")
	fmt.Scan(&menublacksmith)

	switch menublacksmith { // Menu in switch case form, permit to craft elements with blacksmith
	case 1:
		if CheckBlacksmith(perso, "raven feather", 5) {
			if CheckBlacksmith(perso, "wild boar leather", 5) {
				RemoveInventory(perso, "raven feather")
				RemoveInventory(perso, "wild boar leather")
				perso.money -= 5
				AddInventory(perso, "adventurer's hat")
				fmt.Println("You have made an adventurer's hat, it is now added to your inventory")
				time.Sleep(2 * time.Second)
				Blacksmith(perso)
			}
		}

	case 2:
		if CheckBlacksmith(perso, "wolf fur", 5) {
			if CheckBlacksmith(perso, "wolf fur", 5){
				if CheckBlacksmith(perso, "peau de troll", 5) {
					RemoveInventory(perso, "raven feather")
					RemoveInventory(perso, "wild boar leather")
					perso.money -= 5
					AddInventory(perso, "adventurer's tunic")
					fmt.Println("You have crafted an adventurer's tunic, it is now added to your inventory")
					time.Sleep(2 * time.Second)
					Blacksmith(perso)
				}
			}
		}

	case 3:
		if CheckBlacksmith(perso, "wolf fur", 5) {
			if CheckBlacksmith(perso, "wild boar leather", 5) {
				RemoveInventory(perso, "wolf fur")
				RemoveInventory(perso, "wild boar leather")
				perso.money -= 5
				AddInventory(perso, "adventurer's boots")
				fmt.Println("You have crafted Adventurer's Boots, it is now added to your inventory")
				time.Sleep(2 * time.Second)
				Blacksmith(perso)
			}
		}

	case 4:
		Menu(perso)
	}
}

func CheckBlacksmith(perso *Character, item1 string, moneyneeded int) bool { // Function that check if item is in inventory, and if character has sufficent money for craft items
	if CheckItemInventory(perso, item1) {
			if perso.money > moneyneeded {
				return true
			} else { 
				fmt.Println("You don't have enough money to pay for this production!")
				time.Sleep(2 * time.Second)
				Blacksmith(perso)
				return false
			}
	} else {
		fmt.Println("You do not have", item1, "in your inventory!")
		time.Sleep(2 * time.Second)
		Blacksmith(perso)
		return false
	}
}