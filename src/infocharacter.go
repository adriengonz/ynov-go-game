package main

import "fmt"

func DisplayInfo(perso *Character) { // Function that prints specs about the character
	fmt.Println("Nom :", perso.name)
	fmt.Println("Race :", perso.race)
	fmt.Println("Class :", perso.class)
	fmt.Println("Level :", perso.level)
	fmt.Println("Max standard of living :", perso.maxlife)
	fmt.Println("Current standard of living :", perso.currentlife)
	fmt.Println("Usable spell:", perso.skill)
	fmt.Println("Money :", perso.money)
	fmt.Println("Type 0 to return to the previous menu")
	var userinputinfo int
	fmt.Scan(&userinputinfo)
	for userinputinfo != 0 { // While userinput is not 0, repeat the user input
		fmt.Println("Your order was not recognized, type 0 to return to the previous menu")
		fmt.Scan(&userinputinfo)
	}
	Menu(perso)
}

func Dead(perso *Character) { // Function that check if character is dead
	if perso.currentlife == 0 {
		fmt.Println("You are dead")
		perso.currentlife = perso.maxlife / 2
		fmt.Println("You have just been resurrected with half your life points")
	}
}