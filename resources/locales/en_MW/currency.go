package en_MW

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MWK", DisplayName: "", Symbol: "MK"},
	}
}
