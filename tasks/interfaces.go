package tasks

import "fmt"

type Duck struct {
	View   string
	Weight int
}

func (duck Duck) MakeSound() {
	fmt.Println("Quack")
}

func (duck Duck) Walk() {
	fmt.Println("Walk")
}

type Animal interface {
	MakeSound()
	Walk()
}
