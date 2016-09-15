package main

import "fmt"

func main() {
	dog := &Dog{
		Animal: Animal{
			Name: "Rover",
			mean: true,
		},
		BarkStrength: 2,
	}

	cat := &Cat{
		Basics: Animal{
			"Julius",
			true,
		},
		MeowStrength: 5,
	}

	MakeSomeNoise(dog)
	MakeSomeNoise(cat)
}

type Animal struct {
	Name string
	mean bool
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

type AnimalSounder interface {
	MakeNoise()
}

func (animal Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s", sound)
	}

	fmt.Println("")

}

func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Cat) MakeNoise() {
	cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}
