package main

type Equipment struct { // Structure of Equipment
	head  []string
	torso []string
	foot  []string
}

func (stuff *Equipment) Init(head []string, torso []string, foot []string) {
	stuff.head = []string{"Emplacement du casque"}
	stuff.torso = []string{"Emplacement du torse"}
	stuff.foot = []string{"Emplacement des bottes"}
}