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

func (a agent) drive(where string) {
	fmt.Println("WROOOOM!!!", where)
}

func (p person) drive(where string) {
	fmt.Println(".br.r.r.r.", where,"???!!!.r.r.r.rr.rr.rrr...")
}

type Driveable interface {
	drive(where string)
}

func driveMeToMars(driveable Driveable) {
	switch v := driveable.(type) {
	case person:
		v.drive("to Mars???!!! -- no way")
	case agent:
		v.drive("to Mars!!")
	}
}

func main() {
	var p1 person;
	p1 = person{"Ievgen", 31}
	a1 := agent{
		person{
			"Ivan",
			54,
		},
		"James",
		3,
	}
	driveMeToMars(a1)
	driveMeToMars(p1)
}