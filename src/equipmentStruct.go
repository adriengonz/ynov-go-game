package main

type Equipment struct { // Structure of Equipment
	head  []string
	torso []string
	foot  []string
}

func (stuff *Equipment) Init(head []string, torso []string, foot []string) {
	stuff.head = []string{"Headset location"}
	stuff.torso = []string{"Torso location"}
	stuff.foot = []string{"Boot Location"}
}
