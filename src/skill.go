package main

import(
	"fmt"
	"time"
)

func Spellbook(perso *Character, sortbook string) {
	newsortname := sortbook[9:]
	for _, skillverif := range perso.skill { // Check if sort is already in skills slice
		if skillverif == newsortname {
            fmt.Println(newsortname, "est déjà dans votre liste de skills !")
			Menu(perso)
		}
	}
	perso.skill = append(perso.skill, newsortname) // Append sort in skills slice
	RemoveInventory(perso, sortbook)
	fmt.Println(newsortname, "a bien été ajouté à vos skills")
	time.Sleep(2 * time.Second)
	Menu(perso)
}