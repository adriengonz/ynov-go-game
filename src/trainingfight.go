package main

import (
	"fmt"
	"time"
)

func trainingFight(perso *Character, monster *Monster) {
	turn := 1

	fmt.Println("Bienvenue dans le combat d'entraînement!")
	
	for perso.currentlife > 0 && monster.currentlife > 0 {
		if turn%3 == 0 {
			monster.attackpoint *= 2
		}
		fmt.Println("\nTour", turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)

		playerAttack := 5 // Tour du joueur
		fmt.Println("Vous attaquez le monstre et lui infligez", playerAttack, "points de dégâts!")
		monster.currentlife -= playerAttack
		time.Sleep(2 * time.Second)

		if monster.currentlife <= 0 {
			fmt.Println("Vous avez vaincu le monstre! Félicitations!")
			time.Sleep(3 * time.Second)
			break
		}

		// Tour du monstre
		fmt.Println("Le monstre vous attaque et vous inflige", monster.attackpoint, "points de dégâts!")
		perso.currentlife -= monster.attackpoint
		time.Sleep(2 * time.Second)

		if Dead(perso) {
			fmt.Println("You are dead")
			perso.currentlife = perso.maxlife / 2
			fmt.Println("Vous venez de ressuciter avec la moitié de vos points de vies")
			time.Sleep(3 * time.Second)
			Menu(perso, monster)
		}

		turn++
	}

	fmt.Println("Fin du combat d'entraînement.")
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}
