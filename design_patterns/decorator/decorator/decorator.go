package decorator

import "lld/design_patterns/decorator/coffee"

type CoffeeDecorator struct {
    Coffee coffee.Coffee
}

type Milk struct{
    CoffeeDecorator
}

func NewMilk(coffee coffee.Coffee) *Milk {
    return &Milk{
        CoffeeDecorator: CoffeeDecorator{         
            Coffee: coffee,     
        },
    }
}

func (m *Milk) Cost() int {
    return m.Coffee.Cost()+ 5   
}

func (m *Milk) Description() string {
    return m.Coffee.Description() + ", Milk"
}

type WhippedCream struct{
    CoffeeDecorator
}

func NewWhippedCream(coffee coffee.Coffee) *WhippedCream {
    return &WhippedCream{
        CoffeeDecorator: CoffeeDecorator{         
            Coffee: coffee,     
        },
    }
}

func (w *WhippedCream) Cost() int {
    return w.Coffee.Cost()+ 10   
}

func (w *WhippedCream) Description() string {
    return w.Coffee.Description() + ", Whipped Cream"
}

type ChocolateSyrup struct{
    CoffeeDecorator
}

func NewChocolateSyrup(coffee coffee.Coffee) *ChocolateSyrup {
    return &ChocolateSyrup{
        CoffeeDecorator: CoffeeDecorator{         
            Coffee: coffee,
        },
    }
}

func (c *ChocolateSyrup) Cost() int {
    return c.Coffee.Cost()+ 15   
}

func (c *ChocolateSyrup) Description() string {
    return c.Coffee.Description() + ", Chocolate Syrup"
}

type CaramelSyrup struct{
    CoffeeDecorator
}

func NewCaramelSyrup(coffee coffee.Coffee) *CaramelSyrup {
    return &CaramelSyrup{
        CoffeeDecorator: CoffeeDecorator{         
            Coffee: coffee,
        },
    }
}

func (c *CaramelSyrup) Cost() int {
    return c.Coffee.Cost() + 20   
}

func (c *CaramelSyrup) Description() string {
    return c.Coffee.Description() + ", Caramel Syrup"
}