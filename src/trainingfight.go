package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	trainingFight()
}

func trainingFight(perso *Character, monster *Monster) {
	playerHealth := 100
	monsterHealth := 100
	turn := 1

	fmt.Println("Bienvenue dans le combat d'entraînement!")

	for playerHealth > 0 && monsterHealth > 0 {
		fmt.Printf("\nTour %d\n", turn)
		fmt.Printf("Joueur - Points de vie: %d\n", playerHealth)
		fmt.Printf("Monstre - Points de vie: %d\n", monsterHealth)

		// Tour du joueur
		playerAttack := rand.Intn(20) + 1 // Attaque aléatoire entre 1 et 20
		fmt.Printf("Vous attaquez le monstre et lui infligez %d points de dégâts!\n", playerAttack)
		monsterHealth -= playerAttack

		if monsterHealth <= 0 {
			fmt.Println("Vous avez vaincu le monstre! Félicitations!")
			break
		}

		// Tour du monstre
		monsterAttack := rand.Intn(15) + 1 // Attaque aléatoire entre 1 et 15
		fmt.Printf("Le monstre vous attaque et vous inflige %d points de dégâts!\n", monsterAttack)
		playerHealth -= monsterAttack

		if playerHealth <= 0 {
			fmt.Println("Le monstre vous a vaincu. Vous avez perdu le combat.")
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat d'entraînement.")
}
