package en_GM

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GMD", DisplayName: "", Symbol: "D"},
	}
}
