package mgh

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MZN", DisplayName: "", Symbol: "MTn"},
	}
}
