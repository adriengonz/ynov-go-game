package main

import (
	"fmt"
	"os"
)

func Menu(persovar *Character, monsterinstance *Monster) { // Fonction qui affiche l'interface du menu principal
	menuinput := 0
	Clear()
	fmt.Println("Menu:")
	fmt.Println("┌--------------------------------------------┐")
	fmt.Println("| 1- Afficher les informations du personnage |")
	fmt.Println("| 2- Accéder au contenu de l’inventaire      |")
	fmt.Println("| 3- Marchand                                |")
	fmt.Println("| 4- Forgeron                                |")
	fmt.Println("| 5- S'entrainer au combat                   |")
	fmt.Println("| 6- Qui sont ils ?                          |")
	fmt.Println("| 0- Quitter                                 |")
	fmt.Println("└--------------------------------------------┘")
	fmt.Println("Votre choix ?")
	fmt.Scan(&menuinput)
	for menuinput < 0 || menuinput > 6 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
	fmt.Scan(&menuinput)
	}
	switch menuinput { // Code pour sélectionner la fonction voulu
	case 1:
		DisplayInfo(persovar, monsterinstance)
	case 2:
		AccessInventory(persovar, monsterinstance)
	case 3:
		Merchant(persovar, monsterinstance)
	case 4:
		Blacksmith(persovar, monsterinstance)
	case 5:
		TrainingFight(persovar, monsterinstance)
	case 6:
		fmt.Println("Thanos = Josh Brolin","\n","Maximus Decimus Meridius = Russel Crowe")
	case 0:
		os.Exit(0)
	}
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
