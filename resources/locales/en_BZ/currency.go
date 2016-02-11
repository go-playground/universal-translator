package en_BZ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BZD", DisplayName: "", Symbol: "$"},
	}
}
