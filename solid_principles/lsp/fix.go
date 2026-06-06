package main

import (
	"fmt"
)

type Bird interface {
    Name() string
}

type FlyingBird interface {
    Bird
    Fly() error
}

type Sparrow struct{}

func (s *Sparrow) Name() string {
	return "sparrow"
}

func (s *Sparrow) Fly() error {
	fmt.Println("Sparrow is flying")
	return nil
}

type Penguin struct{}

func (p *Penguin) Name() string {
	return "penguin"
}

func MakeBirdFly(b FlyingBird) error {
	return b.Fly()
}

func main() {
	flyingbirds := []FlyingBird{
		&Sparrow{},
	}

	for _, bird := range flyingbirds {
		if err := MakeBirdFly(bird); err != nil {
			fmt.Println("Error:", err)
		}
	}
}