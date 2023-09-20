package main

import "fmt"

func DisplayInfo(perso *Character) { // Function that prints specs about the character
	fmt.Println("Nom :", perso.name)
	fmt.Println("Race :", perso.race)
	fmt.Println("Classe :", perso.class)
	fmt.Println("Niveau :", perso.level)
	fmt.Println("Niveau de vie max :", perso.maxlife)
	fmt.Println("Niveau de vie actuel :", perso.currentlife)
	fmt.Println("Flèches restantes :", perso.arrows)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	var userinputinfo int
	fmt.Scan(&userinputinfo)
	for userinputinfo != 0 { // While userinput is not 0, repeat the user input
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
		fmt.Scan(&userinputinfo)
	}
	Menu(perso)
}

func Dead(perso *Character) { // Function that check if character is dead
	if perso.currentlife == 0 {
		fmt.Println("You are dead")
		perso.currentlife = perso.maxlife / 2
		fmt.Println("Vous venez de ressuciter avec la moitié de vos points de vies")
	}
}