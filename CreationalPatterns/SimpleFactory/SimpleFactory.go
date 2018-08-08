package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s\n", p.Name)
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
	}
}

func NewPersonPointer(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

func main() {
	p1 := NewPerson("John", 12)
	p1.Greet()

	p2 := NewPersonPointer("Arc", 18)
	p2.Greet()
}
