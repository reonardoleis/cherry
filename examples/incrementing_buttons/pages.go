package main

import (
	"github.com/reonardoleis/cherry/pkg/page"
)

func Home() *page.Page {
	navigateButton := NewNavigateButton("other")

	return page.NewPage(NewContainer(navigateButton))
}

func Counter() *page.Page {
	button := NewButton()

	return page.NewPage(NewContainer(button))
}
