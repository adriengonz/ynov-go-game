package main

import (
	"fmt"
	"time"
	"math/rand"
)

func TrainingFight(perso *Character, monster *Monster) { // Combat d'entrainement
	Clear()
	img("./img/fight.jpeg")
	rand.Seed(time.Now().UnixNano()) // Initialiser la génération de nombres aléatoires	
	initiativeplayer := rand.Intn(100-0)
	initiativemonster := rand.Intn(100-0)
	fmt.Println("Bienvenue dans le combat d'entraînement!")
	fmt.Println("Initiatives du combat:\nJoueur:", initiativeplayer, "\nMonstre:", initiativemonster)

	for perso.currentlife > 0 && monster.currentlife > 0 { // Boucle pour le combat 
		fmt.Println("\nTour", perso.turn)
		fmt.Println("Joueur - Points de vie :", perso.currentlife)
		fmt.Println("Monstre - Points de vie :", monster.currentlife)
		fmt.Println("----------------------------------------------")

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
				perso.money += 50 // Ajout d'argent suite a la victoire contre le monstre
				fmt.Println("Vous gagnez 50 pièces d'or !")
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
	monster.attackpoint = 5 // Reinitialisation points d'attaque monstre
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}

func CharTurn(perso *Character, monster *Monster) { // Fonction qui va être appelée au tour du joueur
	attackinput := 0
	fmt.Println("\nQue voulez vous faire ?")
	fmt.Println("1 - attaquer")
	fmt.Println("2 - inventaire")
	fmt.Println("-------------------------")
	fmt.Scan(&attackinput)
	for attackinput < 0 || attackinput > 2 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue")
	fmt.Scan(&attackinput)
	}

	switch attackinput {
		case 1:
			AttackMenu(perso, monster)
		case 2:
			AccessInventoryFight(perso, monster)
		}
}

func AttackMenu(perso *Character, monster *Monster) { // Menu qui s'affiche lors de l'attaque
	attackmenuinput := 0
	fmt.Println("\nQuelle attaque voulez-vous utiliser ?")
	fmt.Println("1 - Attaque classique")
	fmt.Println("2 - Sort")
	fmt.Println("0 - Retour au menu précédent")
	fmt.Println("---------------------------------------")
	fmt.Scan(&attackmenuinput)
	for attackmenuinput < 0 || attackmenuinput > 2 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
	fmt.Scan(&attackmenuinput)
	}

	switch attackmenuinput {
		case 1: // Attaque classique
		playerAttack := 5
		fmt.Println("Vous attaquez le monstre et lui infligez", playerAttack, "points de dégâts!")
		monster.currentlife -= playerAttack
		time.Sleep(2 * time.Second)
		case 2:
			skillinput := 0
			skillname := ""
			fmt.Println(perso.skill) // Affiche les skills disponibles
			fmt.Println("Quel sort voulez-vous utiliser ? Tapez le numéro correspondant au sort (de gauche à droite)")
			fmt.Scan(&skillinput)
			skillname = perso.skill[skillinput-1] // Récupération du nom du skill
			if skillname == "boule de feu" { // Conditions par rapoort au nom du skill, et actions selon celui ci
				if perso.mana > 5 { // Conditions par rapport au mana
					perso.mana -= 5 // Utilisation de mana
					playerAttack := 15 // Définition des points d'attaque par rapport au sort choisi
					fmt.Println("Vous utilisez 5 de mana pour cet attaque")
					time.Sleep(1 * time.Second)
					fmt.Println("Vous attaquez le monstre avec une boule de feu et lui infligez", playerAttack, "points de dégâts!")
					monster.currentlife -= playerAttack
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Vous n'avez pas assez de mana pour utiliser ce sort ! Il vous manque", 5 - perso.mana, "de mana")
					time.Sleep(2 * time.Second)
					AttackMenu(perso, monster)
				}
			} else if skillname == "coup de poing" {
				if perso.mana > 10 { // Conditions par rapport au mana
					perso.mana -= 10 // Utilisation de mana
					playerAttack := 8 // Définition des points d'attaque par rapport au sort choisi
					fmt.Println("Vous utilisez 10 de mana pour cet attaque")
					time.Sleep(1 * time.Second)
					fmt.Println("Vous attaquez le monstre avec un coup de poing et lui infligez", playerAttack, "points de dégâts!")
					monster.currentlife -= playerAttack
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Vous n'avez pas assez de mana pour utiliser ce sort ! Il vous manque", 10 - perso.mana, "de mana")
					time.Sleep(2 * time.Second)
					AttackMenu(perso, monster)
				}
			} else if skillname == "éclair de givre" {
				if perso.mana > 15 { // Conditions par rapport au mana
					perso.mana -= 15 // Utilisation de mana
					playerAttack := 18 // Définition des points d'attaque par rapport au sort choisi
					fmt.Println("Vous utilisez 15 de mana pour cet attaque")
					time.Sleep(1 * time.Second)
					fmt.Println("Vous attaquez le monstre avec un éclair de givre et lui infligez", playerAttack, "points de dégâts!")
					monster.currentlife -= playerAttack
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Vous n'avez pas assez de mana pour utiliser ce sort ! Il vous manque", 15 - perso.mana, "de mana")
					time.Sleep(2 * time.Second)
					AttackMenu(perso, monster)
				}
			}
		case 0: // Retour au menu précédent
			CharTurn(perso, monster)
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

func Exp(perso *Character, monster *Monster) { // Fonction qui ajoute de l'expérience au joueur
	perso.currentExperience += monster.exppoint
	fmt.Println("Vous avez aquéri", monster.exppoint, "points d'éxpérience !")
	time.Sleep(2 * time.Second)
	if perso.currentExperience >= 100 { // Si l'exp est supérieur a 100, passer au niveau supérieur
		perso.currentExperience = 0
		perso.level += 1
		fmt.Println("Vous êtes monté au niveau supérieur ! Vous êtes au niveau", perso.level)
		time.Sleep(2 * time.Second)
	}
}