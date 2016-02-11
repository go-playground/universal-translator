package en_BS

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BSD", DisplayName: "", Symbol: "$"},
	}
}
