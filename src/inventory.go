package main

import (
	"fmt"
	"time"
)

func AccessInventory(perso *Character) { // Function for access inventory content
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	var userinput int
	fmt.Scan(&userinput)
	for userinput != 0 { // While userinput is not 0, repeat the user input
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
		fmt.Scan(&userinput)
	}
	Menu(perso)
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
	if len(perso.inventory) >= 10 {
		return true
	} else {
		return false
	}
}

func CheckItemInventory(perso *Character, itemname string) bool { // Fucntion that check if item is in inventory
	for _, itempicker := range perso.inventory { 
		if itempicker == itemname {
			return true
		}
	}
	return false
}