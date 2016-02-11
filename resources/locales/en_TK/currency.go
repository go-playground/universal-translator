package en_TK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "NZD", DisplayName: "", Symbol: "$"},
	}
}
