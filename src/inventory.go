package main

import (
	"fmt"
)

func accessInventory() {
	fmt.Println(perso1.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	var userinput int
	fmt.Scan(&userinput)
	for userinput != 0 {
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
		fmt.Scan(&userinput)
	}
}
