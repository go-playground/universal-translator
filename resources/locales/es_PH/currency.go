package es_PH

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PHP", DisplayName: "", Symbol: "₱"},
	}
}
