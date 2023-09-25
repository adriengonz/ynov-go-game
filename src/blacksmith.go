package main

import (
	"fmt"
	"time"
)

func Blacksmith(perso *Character, monster *Monster) {
	menublacksmith := 0
	fmt.Println("Bienvenue chez le forgeron")
	fmt.Println("1 - Fabriquer un chapeau de l'aventurier (5 pièces d'or, 1 plume de corbeau et 1 cuir de sanglier)")
	fmt.Println("2 - Fabriquer une tunique de l'aventurier (5 pièces d'or, 2 fourrure de loup et 1 peau de troll)")
	fmt.Println("3 - Fabriquer des bottes de l'aventurier (5 pièces d'or, 1 fourrure de loup et 1 cuir de sanglier)")
	fmt.Println("0 - Revenir au menu principal")
	fmt.Scan(&menublacksmith)

	switch menublacksmith { // Menu in switch case form, permit to craft elements with blacksmith
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

func CheckBlacksmith(perso *Character, monster *Monster, item1 string, moneyneeded int) bool { // Function that check if item is in inventory, and if character has sufficent money for craft items
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
		fmt.Println("Vous n'avez pas de", item1, "dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Blacksmith(perso, monster)
		return false
	}
}