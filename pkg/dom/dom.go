package dom

import (
	"fmt"
	"syscall/js"

	"github.com/reonardoleis/cherry/pkg/manager"
)

var id int32

func GetElementById(id string) js.Value {
	element := js.Global().Get(id)
	return element
}

func UpdateDOM() {
	manager := manager.Instance()

	root := GetElementById("root")

	html := manager.ActivePage().HTML()

	root.Set("innerHTML", html)
}

func CreateFunction(fn func(this js.Value, args []js.Value) any, args ...string) string {
	name := fmt.Sprintf("fn_%d", id)
	jsFn := js.FuncOf(fn)
	js.Global().Set(name, jsFn)
	id++

	name += "("

	for idx, arg := range args {
		name += arg
		if idx != len(args)-1 {
			name += ", "
		}
	}

	name += ")"

	return name
}
