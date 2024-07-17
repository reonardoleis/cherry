package main

import (
	"github.com/reonardoleis/cherry/pkg/state"
)

var countryContext *CountryContext

func init() {
	countryContext = new(CountryContext)

	countryContext.currentCountry = state.SetState("")
	countryContext.countryList = state.SetState("")
}

type CountryContext struct {
	currentCountry *state.State[string]
	countryList    *state.State[string]
}
