package page

import (
	"github.com/reonardoleis/cherry/pkg/component"
)

type Page struct {
	rootComponent component.Component[any]
}

func NewPage(rootComponent component.Component[any]) *Page {
	return &Page{rootComponent: rootComponent}
}

func (p *Page) HTML() string {
	return p.rootComponent.Render()
}
