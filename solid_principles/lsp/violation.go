package main

import (
	"fmt"
)

type Bird interface {
	Fly() error
}

type Sparrow struct{}

func (s *Sparrow) Fly() error {
	fmt.Println("Sparrow is flying")
	return nil
}

type Penguin struct{}

func (p *Penguin) Fly() error {
	return fmt.Errorf("penguins cannot fly")
}

func MakeBirdFly(b Bird) error {
	return b.Fly()
}

func main() {
	birds := []Bird{
		&Sparrow{},
		&Penguin{},
	}

	for _, bird := range birds {
		if err := MakeBirdFly(bird); err != nil {
			fmt.Println("Error:", err)
		}
	}
}