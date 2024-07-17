package main

import (
	"fmt"
	"syscall/js"

	"github.com/reonardoleis/cherry/pkg/component"
	"github.com/reonardoleis/cherry/pkg/dom"
	"github.com/reonardoleis/cherry/pkg/parser"
)

type Button struct {
	component.Base[any]
}

func (b Button) Render() string {
	return "<button>Fetch</button>"
}

type Input struct {
	component.Base[any]
	value string
	fetch component.ComponentFunction
}

func (i *Input) Fetch(js.Value, []js.Value) any {
	go func() {
		query := dom.GetElementById("query").Get("value").String()
		country, err := searchCountry(query)
		if err != nil {
			println(err.Error())
		}

		countryContext.setCountry(country.Capital)

	}()
	return nil
}

func NewInput() Input {
	input := Input{}
	input.fetch = dom.CreateFunction(input.Fetch)

	return input
}

func (i Input) Render() string {
	return fmt.Sprintf("<input id='query' onchange='%s' class='text-black'>", i.fetch)
}

type Div struct {
	component.Base[any]
}

func NewDiv() Div {
	var input = NewInput()
	var div = Div{}
	div.Register(input)

	return div
}

func (d Div) CurrentCountry() string {
	if *countryContext.country != "" {
		return fmt.Sprintf("<h1>Capital: %s</h1>", *countryContext.country)
	}

	return ""
}

func (d Div) Render() string {
	return parser.HTML(d, fmt.Sprintf(`<div class="container h-full w-full bg-zinc-800 text-white flex flex-col items-center justify-center">
    %s
    <h2>Search for countries</h2>
    <.Input></.Input>
  </div>`, d.CurrentCountry()))
}
