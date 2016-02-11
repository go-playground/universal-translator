package es_PE

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PEN", DisplayName: "", Symbol: "S/."},
	}
}
