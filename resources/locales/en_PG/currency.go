package en_PG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "PGK", DisplayName: "", Symbol: "K"},
	}
}
