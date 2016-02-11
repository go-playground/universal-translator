package en_BB

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BBD", DisplayName: "", Symbol: "$"},
	}
}
