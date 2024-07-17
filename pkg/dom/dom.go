package dom

import (
	"fmt"
	"syscall/js"
)

var id int32

func GetElementById(id string) js.Value {
	element := js.Global().Get(id)
	return element
}

func GetElementsByBind(key string) []js.Value {
	document := js.Global().Get("document")
	elements := document.Call("querySelectorAll", fmt.Sprintf(`[bind="%s"]`, key))
	results := make([]js.Value, elements.Length()-1)

	for i := range elements.Length() {
		results = append(results, elements.Index(i))
	}

	return results
}

func CreateFunction(fn func(this js.Value, args []js.Value) any, args ...string) string {
	name := fmt.Sprintf("fn_%d", id)
	jsFn := js.FuncOf(fn)
	js.Global().Set(name, jsFn)
	id++

	name += "("

	for idx, arg := range args {
		name += fmt.Sprintf(`"%s"`, arg)
		if idx != len(args)-1 {
			name += ", "
		}
	}

	name += ")"

	return name
}
