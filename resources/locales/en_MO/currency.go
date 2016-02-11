package en_MO

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MOP", DisplayName: "", Symbol: "MOP$"},
	}
}
