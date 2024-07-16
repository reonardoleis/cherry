package main

import (
	"sync"

	"github.com/reonardoleis/cherry/pkg/dom"
	"github.com/reonardoleis/cherry/pkg/manager"
	"github.com/reonardoleis/cherry/pkg/page"
)

func main() {
	manager := manager.Instance()
	page := page.NewPage(NewButtonContainer())

	manager.RegisterPage(page)

	dom.UpdateDOM()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
