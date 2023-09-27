package main

import(
	"fmt"
	"time"
)

func Spellbook(perso *Character, monster *Monster, sortbook string) {
	if sortbook == "no" {
		fmt.Println("Pas de livre trouvé dans votre inventaire !")
		time.Sleep(2 * time.Second)
		Menu(perso, monster)
	}
	newsortname := sortbook[15:]
	for _, skillverif := range perso.skill { // Vérifie si le skill n'est pas déjà apprit
		if skillverif == newsortname {
            fmt.Println(newsortname, "est déjà dans votre liste de skills !")
			Menu(perso, monster)
		}
	}
	perso.skill = append(perso.skill, newsortname) // Rajoute le skill dans les compétences
	RemoveInventory(perso, sortbook)
	fmt.Println(newsortname, "a bien été ajouté à vos skills")
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}

func SearchBook(perso *Character, monster *Monster) string { // Fonction qui recherche un sort dans un livre de sort
	for _, booksearched := range perso.inventory { // Vérifie si le skill n'est pas déjà apprit
		if booksearched[:15] == "Livre de sort: " {
			return booksearched
		}
	}
	return "no"
}