package ar_SO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SOS", DisplayName: "", Symbol: "S"},
	}
}
