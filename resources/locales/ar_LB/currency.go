package ar_LB

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SDG", DisplayName: "", Symbol: "SDG"},
	}
}
