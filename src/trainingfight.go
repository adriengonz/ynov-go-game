package main

import (
	"fmt"
	"time"
)

func trainingFight(perso *Character, monster *Monster) {
	turn := 1

	fmt.Println("Bienvenue dans le combat d'entraînement!")

	for perso.currentlife > 0 && monster.currentlife > 0 {
		fmt.Println("\nTour", turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)
		if turn%3 == 1 {
			monster.attackpoint = monster.attackpoint*2
		}

		playerAttack := 5 // Tour du joueur
		fmt.Println("Vous attaquez le monstre et lui infligez", playerAttack, "points de dégâts!")
		monster.currentlife -= playerAttack
		time.Sleep(2 * time.Second)

		if monster.currentlife <= 0 {
			fmt.Println("Vous avez vaincu le monstre! Félicitations!")
			break
		}

		monsterAttack := 5 // Tour du monstre
		fmt.Println("Le monstre vous attaque et vous inflige", monsterAttack, "points de dégâts!")
		perso.currentlife -= monsterAttack
		time.Sleep(2 * time.Second)

		if perso.currentlife <= 0 {
			fmt.Println("Le monstre vous a vaincu. Vous avez perdu le combat.")
			Menu(perso, monster)
		}

		turn++
	}

	fmt.Println("Fin du combat d'entraînement.")
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}
