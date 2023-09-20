package main

import (
	"fmt"
	"os"
)

func Merchant(perso *Character) {
	menuinput := 0
	fmt.Println("Bienvenue aventurié")
	fmt.Println("1- Vendre une potion de soin (gratuitement)")
	fmt.Println("2- Vendre une potion de poison (gratuitement)")
	fmt.Println("3- Vendre un livre de sort = boule de feu")
	fmt.Println("4- Revenire dans le menu principale")
	fmt.Scan(&menuinput)

	switch menuinput { // Menu in switch case form, permit to execute functions
	case 1:
	case 2:
	case 3:
	case 4:
		os.Exit(0)
	}
}