package main

import (
	"fmt"
	"time"
)

func Blacksmith(perso *Character, monster *Monster) { // Afficher les options de l'interface forgeron
	menublacksmith := 0
	Clear()
	img("./img/blacksmith.png")
	fmt.Println("Bienvenue chez le forgeron")
	fmt.Println("1 - Fabriquer un chapeau de l'aventurier (5 pièces d'or, 1 plume de corbeau et 1 cuir de sanglier)")
	fmt.Println("2 - Fabriquer une tunique de l'aventurier (5 pièces d'or, 2 fourrure de loup et 1 peau de troll)")
	fmt.Println("3 - Fabriquer des bottes de l'aventurier (5 pièces d'or, 1 fourrure de loup et 1 cuir de sanglier)")
	fmt.Println("0 - Revenir au menu principal")
	fmt.Scan(&menublacksmith)
	for menublacksmith < 0 || menublacksmith > 3 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
	fmt.Scan(&menublacksmith)
	}

	switch menublacksmith { // code pour rajouter l'item craft dans son inventaire, retirer les items nécessaires pour le craft, retirer l'argent du joueur
	case 1:
		if CheckBlacksmith(perso, monster, "plume de corbeau", 5) {
			if CheckBlacksmith(perso, monster, "cuir de sanglier", 5) {
				RemoveInventory(perso, "plume de corbeau")
				RemoveInventory(perso, "cuir de sanglier")
				perso.money -= 5
				perso.equipment.head = []string{"chapeau de l'aventurier"}
				perso.maxlife += 10
				fmt.Println("Vous avez fabriqué un chapeau de l'aventurier, il est désormais ajouté a votre équipement")
				time.Sleep(2 * time.Second)
				Blacksmith(perso, monster)
			}
		}

	case 2:
		if CheckBlacksmith(perso, monster, "fourrure de loup", 5) {
			if CheckBlacksmith(perso, monster, "fourrure de loup", 5){
				if CheckBlacksmith(perso, monster, "peau de troll", 5) {
					RemoveInventory(perso, "plume de corbeau")
					RemoveInventory(perso, "cuir de sanglier")
					perso.money -= 5
					perso.equipment.torso = []string{"tunique de l'aventurier"}
					perso.maxlife += 25
					fmt.Println("Vous avez fabriqué une tunique de l'aventurier, il est désormais ajouté a votre inventaire")
					time.Sleep(2 * time.Second)
					Blacksmith(perso, monster)
				}
			}
		}

	case 3:
		if CheckBlacksmith(perso, monster, "fourrure de loup", 5) {
			if CheckBlacksmith(perso, monster, "cuir de sanglier", 5) {
				RemoveInventory(perso, "fourrure de loup")
				RemoveInventory(perso, "cuir de sanglier")
				perso.money -= 5
				perso.equipment.foot = []string{"bottes de l'aventurier"}
				perso.maxlife += 15
				fmt.Println("Vous avez fabriqué des bottes de l'aventurier, il est désormais ajouté a votre inventaire")
				time.Sleep(2 * time.Second)
				Blacksmith(perso, monster)
			}
		}

	case 0:
		Menu(perso, monster)
	}
}

func CheckBlacksmith(perso *Character, monster *Monster, item1 string, moneyneeded int) bool { // Fonction qui vérifie si l'objet est dans l'inventaire, et si le personnage a assez d'argent pour fabriquer l'objet
	if CheckItemInventory(perso, item1) {
			if perso.money > moneyneeded {
				return true
			} else { 
				fmt.Println("Vous n'avez pas assez d'argent pour payer cette fabrication !")
				time.Sleep(2 * time.Second)
				Blacksmith(perso, monster)
				return false
			}
	} else {
		fmt.Println("Vous n'avez pas de", item1, "dans votre inventaire, allez voir le marchant pour avoir les composants requis !")
		time.Sleep(2 * time.Second)
		Blacksmith(perso, monster)
		return false
	}
}