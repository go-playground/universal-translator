package en_UG

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "UGX", DisplayName: "", Symbol: "USh"},
	}
}
