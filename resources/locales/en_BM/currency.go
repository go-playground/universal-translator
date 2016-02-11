package en_BM

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BMD", DisplayName: "", Symbol: "$"},
	}
}
