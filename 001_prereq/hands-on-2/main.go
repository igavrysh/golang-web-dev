package main

import "fmt"

type person struct {
	name string
	age int
}

type agent struct {
	person person
	nickname string
	numberOfGuns int
}

func (p person)pSpeak() {
	fmt.Println("Hello, my name is:",
		p.name , ". I'm", p.age, " years old")
}

func (a agent)saSpeak() {
	fmt.Println("Pam param...agent", a.nickname,
		"is here")
	for i := 0; i < a.numberOfGuns; i++ {
		fmt.Print("BOOM!!!")
	}
	fmt.Print("\n")
}

func main() {
	var p person;
	p = person{"Ievgen", 31}
	a := agent{
		person {
			"Ivan",
			54,
		},
		"James",
		3,
	}

	p.pSpeak()

	fmt.Println("Agent nickname", a.nickname)

	a.saSpeak()
	a.person.pSpeak()
}