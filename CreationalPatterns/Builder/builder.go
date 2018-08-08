package main

import (
	"strconv"
	"fmt"
)

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BLUE Color = "blue"
	GREEN      = "green"
	RED        = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Car interface {
	Drive() string
	Stop() string
}

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

type carBuilder struct {
	speed  Speed
	color  Color
	wheels Wheels
}

func (cb *carBuilder) Color(color Color) Builder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) Builder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) TopSpeed(speed Speed) Builder {
	cb.speed = speed
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speed,
		wheels:   cb.wheels,
		color:    cb.color,
	}
}

func NewBuilder() Builder {
	return &carBuilder{}
}

type car struct {
	topSpeed Speed
	wheels   Wheels
	color    Color
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.FormatFloat(float64(c.topSpeed), 'f', -1, 64)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}

func main() {
	builder := NewBuilder().Color(BLUE)

	familyCar := builder.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	fmt.Println(familyCar.Drive())
	fmt.Println(familyCar.Stop())

	sportsCar := builder.Wheels(SteelWheels).TopSpeed(150 * KPH).Build()
	fmt.Println(sportsCar.Drive())
	fmt.Println(familyCar.Stop())
}