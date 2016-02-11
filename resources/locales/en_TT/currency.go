package en_TT

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "TTD", DisplayName: "", Symbol: "$"},
	}
}
