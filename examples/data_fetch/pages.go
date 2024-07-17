package main

import "github.com/reonardoleis/cherry/pkg/page"

func Home() *page.Page {
	div := NewDiv()

	return page.NewPage(div)
}
