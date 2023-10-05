package main

import "fmt"

func DisplayInfo(perso *Character, monster *Monster) { // Fonction pour afficher l'interface du joueur
	Clear()
	fmt.Println("Nom :", perso.name)
	fmt.Println("Race :", perso.race)
	fmt.Println("Classe :", perso.class)
	fmt.Println("Niveau :", perso.level)
	fmt.Println("Niveau de vie max :", perso.maxlife)
	fmt.Println("Niveau de vie actuel :", perso.currentlife)
	fmt.Println("Sort utilisable :", perso.skill)
	fmt.Println("Argent :", perso.money)
	fmt.Println("Equipement :", perso.equipment)
	fmt.Println("Mana :", perso.mana)
	fmt.Println("Expérience :", perso.currentExperience, "/100")
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	var userinputinfo int
	fmt.Scan(&userinputinfo)
	for userinputinfo != 0 { // Si l'utilisateur entre un autre caractère que 0, une erreur est renvoyée
		fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
		fmt.Scan(&userinputinfo)
	}
	Menu(perso, monster)
}

func Dead(perso *Character) bool { // Fonction qui vérifie si le joueur est mort ou non
	if perso.currentlife <= 0 {
		return true
	} else {
		return false
	}
}