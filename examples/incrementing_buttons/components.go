package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"syscall/js"

	"github.com/reonardoleis/cherry/pkg/component"
	"github.com/reonardoleis/cherry/pkg/dom"
	"github.com/reonardoleis/cherry/pkg/router"
)

type Button struct {
	incrementFunction string
}

func (b Button) IncrementCounter(js.Value, []js.Value) any {
	counterContext.setCounter(*counterContext.counter + 1)
	return nil
}

func NewButton() Button {
	button := Button{}

	button.incrementFunction = dom.CreateFunction(button.IncrementCounter)

	return button
}

func (b Button) getClass() string {
	if *counterContext.counter%2 == 0 {
		return "bg-orange-300 rounded-md w-[200px] text-zinc-800"
	}

	return "bg-orange-500 rounded-md w-[200px] text-zinc-800"
}

func (b Button) Render() string {
	return `<button class="` + b.getClass() + `" onclick="` + b.incrementFunction + `">` + strconv.Itoa(*counterContext.counter) + `</button>`
}

type NavigateButton struct {
	navigateFunction string
}

func (n NavigateButton) Navigate(v js.Value, args []js.Value) any {
	router.Instance().Navigate(args[0].String())
	return nil
}

func NewNavigateButton(to string) NavigateButton {
	nb := NavigateButton{}

	nb.navigateFunction = dom.CreateFunction(nb.Navigate, to)

	return nb
}

func (n NavigateButton) Render() string {
	return fmt.Sprintf(
		`<button class="bg-orange-300 rounded-md w-[200px] text-zinc-800" onclick='%s'>Go! (current count: %d)</button>`, n.navigateFunction, *counterContext.counter)
}

type Container struct {
	children []component.Component[any]
}

var once sync.Once

func NewContainer(children ...component.Component[any]) Container {
	go once.Do(func() {
		id := rand.Intn(10) + 1
		country, err := FetchCountry(id)
		if err != nil {
			panic(err)
		}

		countryContext.setCountry(*country)
	})

	return Container{children}
}

func (c Container) Render() string {
	content := strings.Repeat("%s", len(c.children))

	for _, child := range c.children {
		content = fmt.Sprintf(content, child.Render())
	}

	return fmt.Sprintf(`<div class="container flex items-center flex-col justify-center w-[100vw] h-[100vh] bg-zinc-800">Country: %s%s</div>`, countryContext.country.Name, content)
}
