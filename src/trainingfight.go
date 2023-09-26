package main

import (
	"fmt"
	"time"
)

func trainingFight(perso *Character, monster *Monster) {
	turn := 1
	attackinput := 0
	for perso.currentlife > 0 && monster.currentlife > 0 {
		if turn%3 == 0 {
			monster.attackpoint *= 2
		}
		fmt.Println("\nTour", turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)

		fmt.Println("Bienvenue dans le combat d'entraînement!")

		fmt.Println("Que voulez vous faire ?") // Tour du joueur
		fmt.Println("1 - attaquer")
		fmt.Println("2 - inventaire")
		fmt.Scan(&attackinput)

		switch attackinput {
		case 1:
		playerAttack := 5
		fmt.Println("Vous attaquez le monstre et lui infligez", playerAttack, "points de dégâts!")
		monster.currentlife -= playerAttack
		time.Sleep(2 * time.Second)
		case 2:
			AccessInventory(perso, monster)
		}

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

func AccessInventoryFight(perso *Character, monster *Monster) { // Function for access inventory content
	menuinventoryfight := 0
	fmt.Println("Voici votre inventaire")
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	fmt.Println("Tapez 1 pour prendre une potion de soin")

	fmt.Scan(&menuinventoryfight)

	switch menuinventoryfight { // Menu in switch case form, permit to execute functions
	case 0:
		trainingFight(perso, monster)
	case 1:
		TakePotFight(perso, monster)
	}
}
