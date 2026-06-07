package main

import (
    "fmt"
    "github.com/aditya-goyal1694/LLD-Prep/design_patterns/abstract_factory/factory"
)

func main() {
    theme := "mac"
    factory, err := factory.NewUIFactory(theme)
    if err != nil {
        fmt.Println("Error creating factory:", err)
        return
    }

    fmt.Println("Factory created successfully")

    button := factory.CreateButton()
    checkbox := factory.CreateCheckbox()
    fmt.Println("Button and checkbox created successfully")

    button.Render()
    checkbox.Render()
    fmt.Println("Button and checkbox rendered successfully")
}