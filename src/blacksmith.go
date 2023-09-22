package main

import (
	"fmt"
	"time"
)

func Blacksmith(perso *Character) {
	menublacksmith := 0
	fmt.Println("Bienvenue chez le forgeron")
	fmt.Println("1 - Fabriquer un chapeau de l'aventurier (5 pièces d'or, 1 plume de corbeau et 1 cuir de sanglier)")
	fmt.Println("2 - Fabriquer une tunique de l'aventurier (5 pièces d'or, 2 fourrure de loup et 1 peau de troll)")
	fmt.Println("3 - Fabriquer des bottes de l'aventurier (5 pièces d'or, 1 fourrure de loup et 1 cuir de sanglier)")
	fmt.Println("4 - Revenir au menu principal")
	fmt.Scan(&menublacksmith)

	switch menublacksmith { // Menu in switch case form, permit to craft elements with blacksmith
	case 1:
		if CheckBlacksmith(perso, "plume de corbeau", "cuir de sanglier", 5) {
			RemoveInventory(perso, "plume de corbeau")
			RemoveInventory(perso, "cuir de sanglier")
			perso.money -= 5
			AddInventory(perso, "chapeau de l'aventurier")
			fmt.Println("Vous avez fabriqué un chapeau de l'aventurier, il est désormais ajouté a votre inventaire")
			time.Sleep(2 * time.Second)
			Blacksmith(perso)
		}
	case 2:
		if CheckBlacksmith(perso, "plume de corbeau", "cuir de sanglier", 5) {
			RemoveInventory(perso, "plume de corbeau")
			RemoveInventory(perso, "cuir de sanglier")
			perso.money -= 5
			AddInventory(perso, "tunique de l’aventurier")
			fmt.Println("Vous avez fabriqué une tunique de l'aventurier, il est désormais ajouté a votre inventaire")
			time.Sleep(2 * time.Second)
			Blacksmith(perso)
		}

	case 3:
		if CheckBlacksmith(perso, "fourrure de loup", "cuir de sanglier", 5) {
			RemoveInventory(perso, "fourrure de loup")
			RemoveInventory(perso, "cuir de sanglier")
			perso.money -= 5
			AddInventory(perso, "bottes de l’aventurier")
			fmt.Println("Vous avez fabriqué des bottes de l'aventurier, il est désormais ajouté a votre inventaire")
			time.Sleep(2 * time.Second)
			Blacksmith(perso)
		}
	case 4:
		
	case 5:
		
	case 6:
		Menu(perso)
	}
}

func CheckBlacksmith(perso *Character, item1 string, item2 string, moneyneeded int) bool {
	if CheckItemInventory(perso, item1) {
		if CheckItemInventory(perso, item2) {
			if perso.money > moneyneeded {
				return true
			} else { 
				fmt.Println("Vous n'avez pas assez d'argent pour payer cette fabrication !")
				time.Sleep(2 * time.Second)
				Blacksmith(perso)
				return false
			}
		} else {
			fmt.Println("Vous n'avez pas de", item2, "dans votre inventaire !")
			time.Sleep(2 * time.Second)
			Blacksmith(perso)
			return false
		}
	} else {
		fmt.Println("Vous n'avez pas de", item1, "dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Blacksmith(perso)
		return false
	}
}