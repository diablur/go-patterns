package main

import "fmt"

type Animal struct {
	name        string
	color       string
	nationality string
}

type AbstractFactory interface {
	CreateAnimal() Animal
}

type CatFactory struct{}

func (*CatFactory) CreateAnimal() Animal {
	return Animal{
		"cat",
		"orange",
		"china",
	}
}

type DogFactory struct{}

func (*DogFactory) CreateAnimal() Animal {
	return Animal{
		"dog",
		"black",
		"usa",
	}
}

func getAnimal(typeAnimal string) Animal {
	var afact AbstractFactory
	switch typeAnimal {
	case "Cat":
		afact = &CatFactory{}
		return afact.CreateAnimal()
	case "Dog":
		afact = &DogFactory{}
		return afact.CreateAnimal()
	}
	return Animal{}
}

func main() {
	a := getAnimal("Cat")
	fmt.Println(a.color)
	fmt.Println(a.nationality)

	b := getAnimal("Dog")
	fmt.Println(b.color)
	fmt.Println(b.nationality)
}
