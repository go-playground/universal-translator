package en_SB

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "SBD", DisplayName: "", Symbol: "$"},
	}
}
