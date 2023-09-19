package main

import "fmt"

func DisplayInfo(perso *Character) {
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
