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
	for _, skillverif := range perso.skill { // Check if sort is already in skills slice
		if skillverif == newsortname {
            fmt.Println(newsortname, "est déjà dans votre liste de skills !")
			Menu(perso, monster)
		}
	}
	perso.skill = append(perso.skill, newsortname) // Append sort in skills slice
	RemoveInventory(perso, sortbook)
	fmt.Println(newsortname, "a bien été ajouté à vos skills")
	time.Sleep(2 * time.Second)
	Menu(perso, monster)
}

func SearchBook(perso *Character, monster *Monster) string { // Function that search a book for spell
	for _, booksearched := range perso.inventory { // Check if sort is already in skills slice
		if booksearched[:15] == "Livre de sort: " {
			return booksearched
		}
	}
	return "no"
}