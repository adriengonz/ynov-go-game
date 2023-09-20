package main

import (
	"fmt"
)

func TakePot(perso *Character) {
	for i, item := range perso.inventory { // Add health to character and remove the potion from the inventory
		if item == "potion de soin" {
			perso.currentlife += 50
            if perso.currentlife > perso.maxlife {
                perso.currentlife = 100
            }
            perso.inventory = append(perso.inventory[:i], perso.inventory[i+1:]...)
			fmt.Println("Vous avez utilisé une potion de soin, votre santé est mainteant de", perso.currentlife)
			Menu(perso)
		}
	}
	fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
	Menu(perso)
}