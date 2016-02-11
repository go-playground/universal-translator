package es_BO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BOB", DisplayName: "", Symbol: "Bs"},
	}
}
