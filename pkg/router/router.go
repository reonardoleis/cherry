package router

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"github.com/reonardoleis/cherry/pkg/dom"
	"github.com/reonardoleis/cherry/pkg/page"
)

var router *Router

type Router struct {
	location   string
	activePage *page.Page
	routes     map[string]*page.Page
}

func init() {
	router = &Router{routes: make(map[string]*page.Page)}

	go func() {
		for {
			location := js.Global().Get("location").Get("href").String()
			location = strings.Split(location, "#")[1]
			location = strings.Replace(location, "#", "", 1)

			if router.location != location {
				router.Handle()
			}
			time.Sleep(time.Millisecond)
		}
	}()
}

func (r *Router) Route(path string, page *page.Page) *Router {
	r.routes[path] = page
	return r
}

func (r *Router) Navigate(to string) {
	if r.routes[to] == nil {
		return
	}

	r.activePage = r.routes[to]

	location := js.Global().Get("location").Get("host").String() + "/#" + to
	protocol := js.Global().Get("location").Get("protocol").String() + "//"
	js.Global().Get("window").Get("location").Set("href", protocol+location)

}

func (r *Router) Handle() {
	location := js.Global().Get("location").Get("href").String()
	location = strings.Split(location, "#")[1]
	location = strings.Replace(location, "#", "", 1)
	r.location = location
	r.activePage = r.routes[location]
	r.UpdateDOM()
}

func (r Router) UpdateDOM() {
	root := dom.GetElementById("root")

	html := r.ActivePage().HTML()

	root.Set("innerHTML", html)
}

func (r Router) UpdateBinds(key string, val any) {
	elements := dom.GetElementsByBind(key)
	for _, element := range elements {
		element.Set("innerHTML", fmt.Sprintf("%v", val))

	}
}

func (r *Router) ActivePage() *page.Page {
	return r.activePage
}

func Instance() *Router {
	return router
}
