package ca_FR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "FRF", DisplayName: "", Symbol: "F"},
	}
}
