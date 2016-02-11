package es_GT

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GTQ", DisplayName: "quetzal", Symbol: "Q"},
	}
}
