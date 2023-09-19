package main

import (
	"fmt"
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
