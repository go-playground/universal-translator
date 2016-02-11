package ar_DJ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DJF", DisplayName: "", Symbol: "Fdj"},
	}
}
