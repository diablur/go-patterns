package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	instance singleton
	once     sync.Once
)

func getInstance() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func main() {
	s1 := getInstance()
	s1["this"] = "rain"
	s2 := getInstance()
	fmt.Println("This is", s2["this"])
}
