package main

import "github.com/reonardoleis/cherry/pkg/router"

func main() {
	router := router.Instance()

	router.Route("", Home())

	router.Handle()
	select {}
}
