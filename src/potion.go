package main

import (
	"fmt"
	"time"
)

func TakePot(perso *Character) { // Add health to character and remove the potion from the inventory
	if RemoveInventory(perso, "healing potion") {
		perso.currentlife += 50
		if perso.currentlife > perso.maxlife {
			perso.currentlife = 100
		}
		fmt.Println("You used a healing potion, your health is now at", perso.currentlife)
		time.Sleep(2 * time.Second)
		Menu(perso)
	} else { // If RemoveInventory return false (so character doesn't have potion in his inventory)
		fmt.Println("You have no potions in your inventory!")
		time.Sleep(2 * time.Second)
		Menu(perso)
	}
}

func PoisonPot(perso *Character) { // Function of poisonpot
	if RemoveInventory(perso, "potion of poison") {
		for i:= 0 ; i < 3 ; i++ {
			perso.currentlife -= 10
			fmt.Println("Damage has been inflicted! You're at", perso.currentlife, "life remaining!")
			time.Sleep(3 * time.Second)
		}
	} else { // If RemoveInventory return false (so character doesn't have potion in his inventory)
		fmt.Println("You have no poison potions in your inventory!")
	}
}