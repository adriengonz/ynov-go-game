package main

import (
	"fmt"
	"time"
)

func TakePot(perso *Character) { // Add health to character and remove the potion from the inventory
	if RemoveInventory(perso, "potion de soin") {
		perso.currentlife += 50
		if perso.currentlife > perso.maxlife {
			perso.currentlife = 100
		}
		fmt.Println("Vous avez utilisé une potion de soin, votre santé est mainteant de", perso.currentlife)
		time.Sleep(2 * time.Second)
		Menu(perso)
	} else {
		fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Menu(perso)
	}
}