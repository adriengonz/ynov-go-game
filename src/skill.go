package main

import(
	"fmt"
	"time"
)

func Spellbook(perso *Character, sortbook string, moneyspend int) {
	newsortname := sortbook[15:]
	checkbook := false
	for _, skillverif := range perso.skill { // Vérifie si le skill n'est pas déjà apprit
		if skillverif == newsortname {
            fmt.Println(newsortname, "est déjà dans votre liste de skills !")
			time.Sleep(3 * time.Second)
			perso.money += moneyspend // Rembourse le joueur si le skill est deja dans la liste de skills
			checkbook = true
		}
	}
	if !checkbook {
		perso.skill = append(perso.skill, newsortname) // Rajoute le skill dans les compétences
		fmt.Println(newsortname, "a bien été ajouté à vos skills")
		time.Sleep(2 * time.Second)
	}
}