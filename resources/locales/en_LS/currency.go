package en_LS

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ZAR", DisplayName: "", Symbol: "R"},
	}
}
