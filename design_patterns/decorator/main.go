package main

import (
	"fmt"
	"lld/design_patterns/decorator/coffee"
	"lld/design_patterns/decorator/decorator"
)
func main() {
    latte := &coffee.Latte{}
    
    latteWithWhippedCream := decorator.NewWhippedCream(latte)
    fmt.Println(latteWithWhippedCream.Description())
    fmt.Printf("The cost is: %d\n", latteWithWhippedCream.Cost())
    
    latteWithWhippedCreamAndCaramelSyrup := decorator.NewCaramelSyrup(latteWithWhippedCream)
    fmt.Println(latteWithWhippedCreamAndCaramelSyrup.Description())
    fmt.Printf("The cost is: %d\n", latteWithWhippedCreamAndCaramelSyrup.Cost())

}











