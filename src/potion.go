package main

import (
	"fmt"
)

func TakePot(perso *Character) {
	var index int = 0
	for i := 1; i < len(perso.inventory); i++ {
		if perso.inventory[i-1] == "potion de soin" {
			index = i
		}
	}
	if index == 0 {
		fmt.Println("Vous n'avez aucune potion dans votre inventaire !")
	} else {
		///RemovePot(perso.inventory, index)
		fmt.Println("tkt")
	}
}

func RemovePot(tab []string, s int) []int {
    return append(tab[:s], tab[s+1:]...)
}