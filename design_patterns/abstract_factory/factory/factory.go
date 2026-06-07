package factory

import (
    "fmt"
    "github.com/aditya-goyal1694/LLD-Prep/design_patterns/abstract_factory/widget"
)


type UIFactory interface {
    CreateButton() widget.Button
    CreateCheckbox() widget.Checkbox
}

type MacUIFactory struct{}

func NewMacUIfactory() *MacUIFactory {
    return &MacUIFactory{}
}

func (m *MacUIFactory) CreateButton() widget.Button {
    return &widget.MacButton{}
}

func (m *MacUIFactory) CreateCheckbox() widget.Checkbox {
    return &widget.MacCheckbox{}
}

type WindowsUIFactory struct{}

func NewWindowsUIfactory() *WindowsUIFactory {
    return &WindowsUIFactory{}
}

func (m *WindowsUIFactory) CreateButton() widget.Button {
    return &widget.WindowsButton{}
}

func (m *WindowsUIFactory) CreateCheckbox() widget.Checkbox {
    return &widget.WindowsCheckbox{}
}

func NewUIFactory(theme string) (UIFactory, error) {
    switch theme {
    case "mac":
        return NewMacUIfactory(), nil
    case "windows":
        return NewWindowsUIfactory(), nil
    default:
        return nil, fmt.Errorf("unsupported theme: %s", theme)
    }
}