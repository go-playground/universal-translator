package ff_MR

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "MRO", DisplayName: "", Symbol: "UM"},
	}
}
