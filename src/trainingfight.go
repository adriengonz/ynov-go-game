package main

import (
	"fmt"
	"time"
	"math/rand"
)

func TrainingFight(perso *Character, monster *Monster) { // Combat d'entrainement
	rand.Seed(time.Now().UnixNano()) // Initialiser la génération de nombres aléatoires	
	initiativeplayer := rand.Intn(100-0)
	initiativemonster := rand.Intn(100-0)
	fmt.Println("Bienvenue dans le combat d'entraînement!")
	fmt.Println("Initiatives du combat:\nJoueur:", initiativeplayer, "\nMonstre:", initiativemonster)

	for perso.currentlife > 0 && monster.currentlife > 0 { // Boucle pour le combat 
		fmt.Println("\nTour", perso.turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)

		if initiativeplayer > initiativemonster { // Si l'initiative du joueur est plus grande, le joueur commence (sinon le monstre commence)
			CharTurn(perso, monster)

		if monster.currentlife <= 0 { // Savoir si le monstre a été vaincu et arrêter le combat
			fmt.Println("Vous avez vaincu le monstre! Félicitations!")
			Exp(perso, monster)
			time.Sleep(3 * time.Second)
			break
		}

		GobelinPattern(perso, monster) // Tour du monstre

		if Dead(perso) { // Savoir si le perosnnage est mort, et ressuciter avec la moitié de ses points de vie max
			fmt.Println("You are dead")
			perso.currentlife = perso.maxlife / 2
			fmt.Println("Vous venez de ressuciter avec la moitié de vos points de vies")
			time.Sleep(3 * time.Second)
			Menu(perso, monster)
		}

		perso.turn++
		} else { // Si le monstre a plus d'initiative que le joueur, le monstre commence
			GobelinPattern(perso, monster)

			if Dead(perso) { // Savoir si le perosnnage est mort, et ressuciter avec la moitié de ses points de vie max
				fmt.Println("You are dead")
				perso.currentlife = perso.maxlife / 2
				fmt.Println("Vous venez de ressuciter avec la moitié de vos points de vies")
				time.Sleep(3 * time.Second)
				Menu(perso, monster)
			}
	
			CharTurn(perso, monster) // Tour du joueur
	
			if monster.currentlife <= 0 { // Savoir si le monstre a été vaincu et arrêter le combat
				fmt.Println("Vous avez vaincu le monstre! Félicitations!")
				time.Sleep(3 * time.Second)
				Exp(perso, monster)
				break
			}
	
			perso.turn++
		}
	}

	fmt.Println("Fin du combat d'entraînement")
	perso.turn = 1 // Reinitialisation tour de combat
	monster.currentlife = monster.maxlife / 2 // Reinitialisation vie monstre
	monster.attackpoint = 5
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}

func AccessInventoryFight(perso *Character, monster *Monster) { // Fonction qui permet d'accéder a l'inventaire pendant un combat (quelques modifications par rapport a l'accès inventaire classique)
	menuinventoryfight := 0
	fmt.Println("Voici votre inventaire")
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	fmt.Println("Tapez 1 pour prendre une potion de soin")

	fmt.Scan(&menuinventoryfight)

	switch menuinventoryfight {
	case 0:
		TrainingFight(perso, monster)
	case 1:
		TakePotFight(perso, monster)
	}
}

func CharTurn(perso *Character, monster *Monster) { // Fonction qui va être appelée au tour du joueur
	attackinput := 0
	fmt.Println("Que voulez vous faire ?")
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
			AccessInventoryFight(perso, monster)
		}
}

func GobelinPattern(perso *Character, monster *Monster) { // Fonction qui va être appelée au tour du monstre
	if perso.turn%3 == 0 {
		monster.attackpoint *= 2
	}
	fmt.Println("Le monstre vous attaque et vous inflige", monster.attackpoint, "points de dégâts!")
	perso.currentlife -= monster.attackpoint
	time.Sleep(2 * time.Second)
}

func Exp(perso *Character, monster *Monster) {
	perso.currentExperience += monster.exppoint
	fmt.Println("Vous avez aquéri", monster.exppoint, "points d'éxpérience !")
	time.Sleep(2 * time.Second)
	if perso.currentExperience >= 100 {
		perso.currentExperience = 0
		perso.level += 1
		fmt.Println("Vous êtes monté au niveau supérieur ! Vous êtes au niveau", perso.level)
		time.Sleep(2 * time.Second)
	}
}