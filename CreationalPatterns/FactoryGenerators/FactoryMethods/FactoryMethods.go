package main

import "fmt"

type Animal struct {
	species string
	age     int
}

type AnimalHouse struct {
	name         string
	sizeInMeters int
}

type AnimalFactory struct {
	species   string
	houseName string
}

func (af AnimalFactory) NewAnimal(age int) Animal {
	return Animal{
		species: af.species,
		age:     age,
	}
}

func (af AnimalFactory) NewHouse(sizeInMeters int) AnimalHouse {
	return AnimalHouse{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

func main() {
	dogFactory := AnimalFactory{
		species:   "dog",
		houseName: "kennel",
	}
	dog := dogFactory.NewAnimal(2)
	fmt.Println(dog.age)
	kennel := dogFactory.NewHouse(3)
	fmt.Println(kennel.sizeInMeters)

	horseFactory := AnimalFactory{
		species:   "horse",
		houseName: "stable",
	}
	horse := horseFactory.NewAnimal(12)
	fmt.Println(horse.age)
	stable := horseFactory.NewHouse(30)
	fmt.Println(stable.name)
}
