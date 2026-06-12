package coffee

type Coffee interface {
    Cost() int
    Description() string
}

type Latte struct{}

func (l *Latte) Cost() int {
    return 10   
}

func (l *Latte) Description() string {
    return "Latte"
}

type Espresso struct{}

func (e *Espresso) Cost() int {
    return 15   
}

func (e *Espresso) Description() string {
    return "Espresso"
}

type Cappuccino struct{}

func (c *Cappuccino) Cost() int {
    return 20   
}

func (c *Cappuccino) Description() string {
    return "Cappuccino"
}
