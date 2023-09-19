package main

import "fmt"

func DisplayInfo(perso *Personne) {
	fmt.Println("Nom :", perso.name)
	fmt.Println("Race :", perso.race)
	fmt.Println("Classe :", perso.class)
	fmt.Println("Niveau :", perso.level)
	fmt.Println("Niveau de vie max :", perso.maxlife)
	fmt.Println("Niveau de vie actuel :", perso.currentlife)
}
