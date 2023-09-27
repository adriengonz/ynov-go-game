package main

import (
	"fmt"
	"time"
)

func TakePot(perso *Character, monster *Monster) { // Rajoute les points de vie que la potion doit donner et supprime l'item
	if RemoveInventory(perso, "potion de soin") {
		perso.currentlife += 50
		if perso.currentlife > perso.maxlife {
			perso.currentlife = 100
		}
		fmt.Println("Vous avez utilisé une potion de soin, votre santé est mainteant de", perso.currentlife)
		time.Sleep(2 * time.Second)
		Menu(perso, monster)
	} else { // Si RemoveInventory renvoie <false> alors ça veut dire qu'on a pas de potion de soin dans l'inventaire
		fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Menu(perso, monster)
	}
}

func PoisonPot(perso *Character, monster *Monster) { // La fonction de la poition de poison
	if RemoveInventory(perso, "potion de poison") {
		for i:= 0 ; i < 3 ; i++ {
			perso.currentlife -= 10
			fmt.Println("Des degats ont été infligés ! L'ennemi est à", monster.currentlife, "de vie !")
			time.Sleep(3 * time.Second)
		}
	} else { // Si RemoveInventory renvoie <false> alors ça veut dire qu'on a pas de potion de poison dans l'inventaire
		fmt.Println("Vous n'avez aucune potion de poison dans votre inventaire !")
	}
}

func TakePotFight(perso *Character, monster *Monster) { // Rajoute les points de vie que la potion doit donner et supprime l'item
	if RemoveInventory(perso, "potion de soin") {
		perso.currentlife += 50
		if perso.currentlife > perso.maxlife {
			perso.currentlife = perso.maxlife
		}
		fmt.Println("Vous avez utilisé une potion de soin, votre santé est mainteant de", perso.currentlife)
		time.Sleep(2 * time.Second)
		TrainingFight(perso, monster)
	} else { // Si RemoveInventory renvoie <false> alors ça veut dire qu'on a pas de potion de soin dans l'inventaire
		fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
		time.Sleep(2 * time.Second)
		TrainingFight(perso, monster)
	}
}