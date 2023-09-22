package main

import (
	"fmt"
	"time"
)

func AccessInventory(perso *Character) { // Function for access inventory content
	menuinventory := 0
	fmt.Println("Here is your inventory")
	fmt.Println(perso.inventory)
	fmt.Println("Type 0 to return to the previous menu")
	fmt.Println("Type 1 to take a healing potion")
	fmt.Println("Type 2 to learn a spell")
	fmt.Scan(&menuinventory)

	switch menuinventory { // Menu in switch case form, permit to execute functions
	case 0:
		Menu(perso)
	case 1:
		TakePot(perso)
	case 2:
		Spellbook(perso, "fireball book")
	}
	var userinput int
	fmt.Scan(&userinput)
	for userinput != 0 { // While userinput is not 0, repeat the user input
		fmt.Println("Your order was not recognized, type 0 to return to the previous menu")
		fmt.Scan(&userinput)
	}
	Menu(perso)
}

func AddInventory(perso *Character, item string) { // Function that add the idem variable to player inventory
	if LimitItem(perso) {
		fmt.Println("The item cannot be added to your inventory because it is full!")
	} else {
		perso.inventory = append(perso.inventory, item)
		fmt.Println(item, "has been added to your inventory")
		time.Sleep(2 * time.Second)
	}
}

func RemoveInventory(perso *Character, itemname string) bool { // Function that remove the item variable from the player inventory
	for i, itempicker := range perso.inventory { 
		if itempicker == itemname {
            perso.inventory = append(perso.inventory[:i], perso.inventory[i+1:]...)
			return true
		}
	}
	return false
}

func LimitItem(perso *Character) bool { // Function that check if limit of inventory is reached
	if len(perso.inventory) >= 10 {
		return true
	} else {
		return false
	}
}

func CheckItemInventory(perso *Character, itemname string) bool { // Function that check if item is in inventory
	for _, itempicker := range perso.inventory { 
		if itempicker == itemname {
			return true
		}
	}
	return false
}