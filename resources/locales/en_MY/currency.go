package en_MY

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MYR", DisplayName: "", Symbol: "RM"},
	}
}
