package main

import "fmt"

type Animal interface {
	Greet()
}

type cat struct{}

func (*cat) Greet()  {
	fmt.Println("I am cat")
}

func newCat() Animal {
	return &cat{}
}

type dog struct {}

func (*dog) Greet()  {
	fmt.Println("I am dog")
}

func newDog() Animal {
	return &dog{}
}

type person struct {}

func (*person) Greet()  {
	fmt.Println("I am person")
}

func newPerson() Animal {
	return &person{}
}

type AnimalType int

const (
	CAT AnimalType = 1 << iota
	DOG
	PERSON
)

func NewAnimal(t AnimalType) Animal {
	switch t {
	case CAT:
		return newCat()
	case DOG:
		return newDog()
	default:
		return newPerson()
	}
}

func main()  {
	c := NewAnimal(CAT)
	c.Greet()

	d := NewAnimal(DOG)
	d.Greet()

	p := NewAnimal(PERSON)
	p.Greet()
}

