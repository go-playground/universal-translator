package en_KN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "XCD", DisplayName: "", Symbol: "$"},
	}
}
