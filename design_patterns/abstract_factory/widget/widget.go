package main

import (
    "fmt"
)

type Widget interface {
    Render()
}

type Button interface{
    Widget
}

type MacButton struct{}

func (m *MacButton) Render() {
    fmt.Printf("Mac button is rendering.")
}

type WindowsButton struct{}

func (w *WindowsButton) Render() {
    fmt.Printf("Windows button is rendering.")
}

type Checkbox interface{
    Widget
}

type MacCheckbox struct{}

func (m *MacCheckbox) Render() {
    fmt.Printf("Mac checkbox is rendering.")
}

type WindowsCheckbox struct{}

func (w *WindowsCheckbox) Render() {
    fmt.Printf("Windows checkbox is rendering.")
}