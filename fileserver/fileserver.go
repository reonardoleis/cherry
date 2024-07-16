package main

import "net/http"

func main() {
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./dist"))); err != nil {
		panic(err)
	}
}
