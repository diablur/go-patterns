package main

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s\n", p.name)
}

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age: age,
	}
}

func NewPersonPointer(name string, age int) Person {
	return &person{
		name: name,
		age: age,
	}
}

func main()  {
	p1 := NewPerson("Smith", 20)
	p1.Greet()
	p2 := NewPersonPointer("Pointer", 16)
	p2.Greet()
}

