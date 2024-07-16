package main

import "github.com/reonardoleis/cherry/pkg/state"

var counterContext *CounterContext
var countryContext *CountryContext

type CounterContext struct {
	counter    *int
	setCounter state.SetFunc[int]
}

type CountryContext struct {
	country    *Country
	setCountry state.SetFunc[Country]
}

func init() {
	counter, setCounter := state.SetState(0)
	counterContext = &CounterContext{counter, setCounter}

	country, setCountry := state.SetState(Country{})
	countryContext = &CountryContext{country, setCountry}
}
