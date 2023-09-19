package main

import (
	"fmt"
)

func AccessInventory() { // Function for access inventory content
	fmt.Println(perso1.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent, ou tapez 1 pour utiliser une potion de vie")
	var userinput int
	fmt.Scan(&userinput)
	for userinput != 0 || userinput != 1 { // While userinput is not 0, repeat the user input
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent ou taper 1 pour utiliser une potion de vie")
		fmt.Scan(&userinput)
	}
	if userinput == 1 {
		TakePot()
	} else {
		Menu()
	}
}
