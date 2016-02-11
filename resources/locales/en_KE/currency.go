package en_KE

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "KES", DisplayName: "", Symbol: "Ksh"},
	}
}
