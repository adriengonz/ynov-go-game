package main

import (
	"fmt"
	"time"
)

func AccessInventory(perso *Character, monster *Monster) { // Function for access inventory content
	menuinventory := 0
	Clear()
	fmt.Println("Voici votre inventaire")
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	fmt.Println("Tapez 1 pour prendre une potion de soin")
	fmt.Println("Tapez 2 pour apprendre un sort")
	fmt.Scan(&menuinventory)

	switch menuinventory { // Menu in switch case form, permit to execute functions
	case 0:
		Menu(perso, monster)
	case 1:
		TakePot(perso, monster)
	case 2:
		book := SearchBook(perso, monster)
		Spellbook(perso, monster, book)
	}
	var userinput int
	fmt.Scan(&userinput)
	for userinput != 0 { // While userinput is not 0, repeat the user input
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
		fmt.Scan(&userinput)
	}
	Menu(perso, monster)
}

func AddInventory(perso *Character, item string) { // Function that add the idem variable to player inventory
	if LimitItem(perso) {
		fmt.Println("L'item ne peut pas être ajouté a votre inventaire car il est plein !")
	} else {
		perso.inventory = append(perso.inventory, item)
		fmt.Println(item, "a bien été ajoutée à votre inventaire")
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
	if len(perso.inventory) >= perso.limitInventory {
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

func UpgradeInventorySlot(perso *Character) {
	counter := 0
	if counter <= 2 {
		perso.limitInventory += 10
		fmt.Println("Votre pouvez stocker dès a présent 10 éléments de plus dans votre invetaire !")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Vous ne pouvez pas améliorer plus de 3 fois votre inventaire !")
		time.Sleep(2 * time.Second)
	}
}