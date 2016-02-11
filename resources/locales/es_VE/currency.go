package es_VE

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "VEF", DisplayName: "", Symbol: "Bs."},
	}
}
