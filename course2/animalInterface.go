package course2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type AnimalStructure struct {
	food  string
	move  string
	sound string
}

func (t *AnimalStructure) Eat()   { fmt.Printf("%s\n", t.food) }
func (t *AnimalStructure) Move()  { fmt.Printf("%s\n", t.move) }
func (t *AnimalStructure) Speak() { fmt.Printf("%s\n", t.sound) }

func GetAnimal(t AnimalInterface, method *string) {
	switch *method {
	case "eat":
		t.Eat()
	case "move":
		t.Move()
	case "speak":
		t.Speak()
	}
}

func CreateAnimal(name, animalType *string, data *map[string]AnimalStructure) {
	(*data)[*name] = (*data)[*animalType]
}

func SetupAnimal() {
	animals := map[string]AnimalStructure{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	var command, input1, input2 *string
	fmt.Println("Use commands newanimal/query to interact. Available queries: speak, eat, move")
	fmt.Println("Create new animals using <newanimal> <name> <type>. Available types: cow, bird, snake")
	for {
		fmt.Printf("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		if len(input) != 3 {
			fmt.Println("Invalid input provided!")
		} else {
			command, input1, input2 = &input[0], &input[1], &input[2]
			switch *command {
			case "newanimal":
				CreateAnimal(input1, input2, &animals)
				fmt.Println("Created it!")
			case "query":
				animal, ok := animals[*input1]
				if ok {
					GetAnimal(&animal, input2)
				}
			}
		}
	}
}
