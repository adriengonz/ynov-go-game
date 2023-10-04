package main

import (
	"fmt"
	"time"
)

func AccessInventory(perso *Character, monster *Monster) { // Fonction pour afficher l'interface de l'inventaire
	menuinventory := 0
	Clear()
	img("./img/backpack.png")
	fmt.Println("Voici votre inventaire")
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	fmt.Println("Tapez 1 pour prendre une potion de soin")
	fmt.Println("Tapez 2 pour prendre une potion de mana (rajoute 25 de mana)")
	fmt.Scan(&menuinventory)
	for menuinventory < 0 || menuinventory > 2 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
	fmt.Scan(&menuinventory)
}

	switch menuinventory { // Menu pour choisir l'action du joueur
	case 0:
		Menu(perso, monster)
	case 1:
		TakePot(perso, monster)
	case 2:
	TakePotMana(perso, monster)
	}
	Menu(perso, monster)
}

func AddInventory(perso *Character, item string) { // Fonction qui rajoute l'item dans l'inventaire
	if LimitItem(perso) {
		fmt.Println("L'item ne peut pas être ajouté a votre inventaire car il est plein !")
	} else {
		perso.inventory = append(perso.inventory, item)
		fmt.Println(item, "a bien été ajoutée à votre inventaire")
		time.Sleep(2 * time.Second)
	}
}

func RemoveInventory(perso *Character, itemname string) bool { // Fonction qui retire l'objet de l'inventaire une fois utilisé
	for i, itempicker := range perso.inventory { 
		if itempicker == itemname {
            perso.inventory = append(perso.inventory[:i], perso.inventory[i+1:]...)
			return true
		}
	}
	return false
}

func LimitItem(perso *Character) bool { // Fonction qui vérifie la limite de l'inventaire
	if len(perso.inventory) >= perso.limitInventory {
		return true
	} else {
		return false
	}
}

func CheckItemInventory(perso *Character, itemname string) bool { // Focntion qui vérifie si l'objet est dans l'inventaire
	for _, itempicker := range perso.inventory { 
		if itempicker == itemname {
			return true
		}
	}
	return false
}

func UpgradeInventorySlot(perso *Character) { // Fonction qui rajoute des places dans l'inventaire si on à l'objet attitré
	counter := 0
	if counter <= 2 {
		perso.limitInventory += 10
		fmt.Println("Votre pouvez stocker dès a présent 10 éléments de plus dans votre invetaire !")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Vous ne pouvez pas améliorer plus de 3 fois votre inventaire !")
		time.Sleep(2 * time.Second)
	}
}

func AccessInventoryFight(perso *Character, monster *Monster) { // Fonction qui permet d'accéder a l'inventaire pendant un combat (quelques modifications par rapport a l'accès inventaire classique)
	menuinventoryfight := 0
	fmt.Println("Voici votre inventaire")
	fmt.Println(perso.inventory)
	fmt.Println("Tapez 0 pour revenir au menu précedent")
	fmt.Println("Tapez 1 pour prendre une potion de soin")
	fmt.Println("Tapez 2 pour prendre une potion de mana")
	fmt.Println("Tapez 3 pour utiliser une potion de poison")
	fmt.Scan(&menuinventoryfight)
	for menuinventoryfight < 0 || menuinventoryfight > 3 { // Si l'utilisateur entre un caractère plus petit ou plus grand, une erreur est renvoyée
	fmt.Println("Votre commande n'a pas été reconnue, tapez 0 pour revenir au menu précédent")
	fmt.Scan(&menuinventoryfight)
	}

	switch menuinventoryfight {
	case 0:
		TrainingFight(perso, monster)
	case 1:
		TakePotFight(perso, monster)
	case 2:
		TakePotManaFight(perso, monster)
	case 3:
		PoisonPot(perso, monster)
	}
}