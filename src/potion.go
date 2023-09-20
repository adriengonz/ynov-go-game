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
	} else { // If RemoveInventory return false (so character doesn't have potion in his inventory)
		fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Menu(perso)
	}
}

func PoisonPot(perso *Character) { // Function of poisonpot
	if RemoveInventory(perso, "potion de poison") {
		for i:= 0 ; i < 3 ; i++ {
			perso.currentlife -= 10
			fmt.Println("Des degats ont été infligés ! Vous êtes à", perso.currentlife, "de vie restante !")
			time.Sleep(3 * time.Second)
		}
	} else { // If RemoveInventory return false (so character doesn't have potion in his inventory)
		fmt.Println("Vous n'avez aucune potion de poison dans votre inventaire !")
	}
}