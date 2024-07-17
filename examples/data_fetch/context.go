package main

import "github.com/reonardoleis/cherry/pkg/state"

var countryContext *CountryContext

func init() {
	countryContext = new(CountryContext)
	countryContext.country, countryContext.setCountry = state.SetState("")
}

type CountryContext struct {
	country    *string
	setCountry state.SetFunc[string]
}
