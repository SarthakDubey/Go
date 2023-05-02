package course2

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (data *Animal) Eat() {
	fmt.Println(data.food)
}

func (data *Animal) Move() {
	fmt.Println(data.locomotion)
}

func (data *Animal) Speak() {
	fmt.Println(data.noise)
}

func AnimalTable() {
	data := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	fmt.Println("Enter data that you need to access: <Animal> <eat/move/speak> eg: Cow Move or cow Move")
	var d1, d2 string
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&d1, &d2)
		if err != nil {
			panic(err)
		}
		animal := data[strings.ToLower(d1)]
		choice := strings.ToLower(d2)
		switch choice {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
