package en_TO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "TOP", DisplayName: "", Symbol: "T$"},
	}
}
