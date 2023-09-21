package main

import (
	"fmt"
)

func Blacksmith(perso *Character) {
	menublacksmith := 0
	fmt.Scan(&menublacksmith)
	fmt.Println("Bienvenue chez le forgeron")
	fmt.Println("1 - Fabriquer un chapeau de l'aventurier (5 pièces d'or, 1 plume de corbeau et 1 cuir de sanglier)")
	fmt.Println("2 - Fabriquer une tunique de l'aventurier (5 pièces d'or, 2 fourrure de loup et 1 peau de troll)")
	fmt.Println("3 - Fabriquer des bottes de l'aventurier (5 pièces d'or, 1 fourrure de loup et 1 cuir de sanglier)")
	fmt.Println("4 - Revenir au menu principal")

	switch menublacksmith { // Menu in switch case form, permit to craft elements with blacksmith
	case 1:
		
	case 2:
		
	case 3:
		
	case 4:
		
	case 5:
		
	case 6:
		Menu(perso)
	}
}