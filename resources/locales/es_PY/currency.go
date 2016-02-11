package es_PY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PYG", DisplayName: "", Symbol: "Gs."},
	}
}
