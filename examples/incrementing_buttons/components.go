package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/reonardoleis/cherry/pkg/dom"
	"github.com/reonardoleis/cherry/pkg/state"
)

type Button struct {
	count             *int
	setCount          state.SetFunc[int]
	incrementFunction string
}

func (b Button) IncrementCounter(js.Value, []js.Value) any {
	b.setCount(*b.count + 1)
	return nil
}

func NewButton() Button {
	count, setCount := state.SetState("count", 0)
	button := Button{count, setCount, ""}

	button.incrementFunction = dom.CreateFunction(button.IncrementCounter)

	return button
}

func (b Button) Render() string {
	return `<button onclick="` + b.incrementFunction + `">` + strconv.Itoa(*b.count) + `</button>`
}

type ButtonContainer struct {
	firstButton  Button
	secondButton Button
}

func NewButtonContainer() ButtonContainer {
	button1 := NewButton()
	button2 := NewButton()

	return ButtonContainer{button1, button2}
}

func (b ButtonContainer) Render() string {
	return fmt.Sprintf(`
    <div>%s<br>%s</div>
    `, b.firstButton.Render(), b.secondButton.Render())
}
