package es_CR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CRC", DisplayName: "", Symbol: "₡"},
	}
}
