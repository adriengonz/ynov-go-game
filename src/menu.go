package main

import (
	"fmt"
	"os"
)

func Menu(persovar *Character) { // Function that prints the menu of game, who takes in argument "persovar" in pointer, who appoint the character structure
	menuinput := 0
	fmt.Println("Menu:")
	fmt.Println("1- Show character information")
	fmt.Println("2- Access inventory content")
	fmt.Println("3- Merchant")
	fmt.Println("4- Blacksmith")
	fmt.Println("5- Leave")
	fmt.Println("Your choice ?")
	fmt.Scan(&menuinput)

	switch menuinput { // Menu in switch case form, permit to execute functions
	case 1:
		DisplayInfo(persovar)
	case 2:
		AccessInventory(persovar)
	case 3:
		Merchant(persovar)
	case 4:
		Blacksmith(persovar)
	case 5:
		os.Exit(0)
	}
}


/*func Menu(perso *Character) {
    // Ajoutez ici la logique de votre menu
    fmt.Println("Menu du jeu")
    fmt.Println("Nom du personnage:", perso.name)
    fmt.Println("Race du personnage:", perso.race)
    fmt.Println("Classe du personnage:", perso.class)
    // Affichez les compétences du personnage
    fmt.Println("Compétences du personnage:")
    for _, skill := range perso.skills {
        fmt.Println(skill)
    }
    // Affichez d'autres informations sur le personnage
}*/