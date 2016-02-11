package es_NI

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "NIO", DisplayName: "", Symbol: "C$"},
	}
}
