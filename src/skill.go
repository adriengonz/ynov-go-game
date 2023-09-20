package main

import(
	"fmt"
	"time"
)

func Spellbook(perso *Character, sort string) {
	for _, skillverif := range perso.skill {
		if skillverif == sort {
            fmt.Println(sort, "est déjà dans votre liste de skills !")
			Menu(perso)
		}
	}
	perso.skill = append(perso.skill, sort)
	fmt.Println(sort, "a bien été ajouté à vos skills")
	time.Sleep(2 * time.Second)
}