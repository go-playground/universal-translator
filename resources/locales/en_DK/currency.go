package en_DK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "DKK", DisplayName: "", Symbol: "kr."},
	}
}
