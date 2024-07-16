# Cherry

<img align="right" width="159px" src="https://i.imgur.com/qz0AVzM.png">

Cherry is a minimalist reactive framework for building frontends in [Go](https://go.dev/). Inspired by React and leveraging the power of WebAssembly (WASM), Cherry allows you to create highly interactive and efficient web applications with Go. While still a work in progress, Cherry aims to combine the simplicity and productivity of modern frontend development with the performance benefits of Go. If you're looking for a fresh approach to frontend development in Go, keep an eye on Cherry as it evolves.
<br/><br/>

## Example: [Incrementing Buttons](https://github.com/reonardoleis/cherry/tree/main/examples/incrementing_buttons) 
```go
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
	return fmt.Sprintf("<button onclick='%s'>%d</button>", b.incrementFunction, *b.count)
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
	return fmt.Sprintf("<div>%s<br>%s</div>", b.firstButton.Render(), b.secondButton.Render())
}
```

### Result
[Video](https://github.com/user-attachments/assets/ffa0b2ae-3dc0-468f-be65-178d151a07b6)

