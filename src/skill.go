package main

import(
	"fmt"
	"time"
)

func Spellbook(perso *Character, sortbook string) {
	newsortname := sortbook[9:]
	for _, skillverif := range perso.skill { // Check if sort is already in skills slice
		if skillverif == newsortname {
            fmt.Println(newsortname, "is already in your list of skills!")
			Menu(perso)
		}
	}
	perso.skill = append(perso.skill, newsortname) // Append sort in skills slice
	RemoveInventory(perso, sortbook)
	fmt.Println(newsortname, "has been added to your skills")
	time.Sleep(2 * time.Second)
	Menu(perso)
}