# Cherry

<img align="right" width="159px" src="https://i.imgur.com/qz0AVzM.png">

Cherry is a minimalist reactive toy framework (**made as a toy project**) for building frontends in [Go](https://go.dev/). Inspired by React and leveraging the power of WebAssembly (WASM), Cherry allows you to create interactive web applications with Go. 
<br/><br/>

## Example: [Live data fetching](https://github.com/reonardoleis/cherry/tree/main/examples/data_fetch) 
```go
type Button struct {
	component.Base[any]
	text        string
	addFunction component.ComponentFunction
}

func (b Button) Add(js.Value, []js.Value) any {
	countryContext.countryList.Set(countryContext.countryList.Get() + fmt.Sprintf("<li>%s</li>", countryContext.currentCountry.Get()))
	return nil
}

func NewButton(text string) Button {
	button := Button{text: text}
	button.addFunction = dom.CreateFunction(button.Add)
	return button
}

func (b Button) Render() string {
	return fmt.Sprintf("<button onclick='%s' class='rounded-md w-[100px] text-black bg-yellow-500'>%s</button>", b.addFunction, b.text)
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

		countryContext.currentCountry.Set(country.Capital)

	}()
	return nil
}

func NewInput() Input {
	input := Input{}
	input.fetch = dom.CreateFunction(input.Fetch)

	return input
}

func (i Input) Render() string {
	return fmt.Sprintf("<input id='query' onkeyup='%s' class='text-black'>", i.fetch)
}

type Div struct {
	component.Base[any]
}

func NewDiv() Div {
	var input = NewInput()
	var button = NewButton("Add")
	var div = Div{}
	div.Register(input)
	div.Register(button)
	countryContext.currentCountry.Bind("currentCountry")
	countryContext.countryList.Bind("countryList")

	return div
}

func (d Div) Render() string {
	return parser.HTML(d, fmt.Sprintf(`<div class="container h-full w-full bg-zinc-800 text-white flex flex-col items-center justify-center">
    <h2>Search for countries</h2>
    <.Input></.Input>
    <div>Current: <state bind="currentCountry"></state></div>
    <.Button></.Button>
    <br>
    <br>
    <ul bind="countryList">
    </ul>
  </div>`))
}
```

### Result
[Video](https://github.com/user-attachments/assets/f119d384-91d7-432e-8c3e-c5f82fa94e75)

