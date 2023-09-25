package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trainingFight(perso *Character, monster *Monster) {
	turn := 1

	fmt.Println("Bienvenue dans le combat d'entraînement!")

	for perso.currentlife > 0 && monster.currentlife > 0 {
		fmt.Println("\nTour %d\n", turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)

		playerAttack := rand.Intn(20) + 1 // Tour du joueur
		fmt.Println("Vous attaquez le monstre et lui infligez %d points de dégâts!\n", playerAttack)
		monster.currentlife -= playerAttack

		if monster.currentlife <= 0 {
			fmt.Println("Vous avez vaincu le monstre! Félicitations!")
			break
		}

		monsterAttack := rand.Intn(15) + 1 // Tour du monstre
		fmt.Println("Le monstre vous attaque et vous inflige %d points de dégâts!\n", monsterAttack)
		perso.currentlife -= monsterAttack

		if perso.currentlife <= 0 {
			fmt.Println("Le monstre vous a vaincu. Vous avez perdu le combat.")
			Menu(perso, monster)
		}

		turn++
	}

	fmt.Println("Fin du combat d'entraînement.")
}
